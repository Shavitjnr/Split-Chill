<template>
    <v-card class="elevation-2 rounded-lg">
        <!-- Calendar Toolbar -->
        <v-toolbar flat color="transparent" class="px-4 pt-2">
            <v-btn variant="outlined" color="primary" class="me-4" @click="setToday">
                {{ tt('Today') }}
            </v-btn>
            <v-btn icon variant="text" color="grey-darken-2" @click="prevMonth">
                <v-icon :icon="mdiChevronLeft"></v-icon>
            </v-btn>
            <v-btn icon variant="text" color="grey-darken-2" @click="nextMonth">
                <v-icon :icon="mdiChevronRight"></v-icon>
            </v-btn>
            
            <v-toolbar-title class="text-h6 font-weight-bold ms-4">
                {{ currentMonthName }} {{ currentYear }}
            </v-toolbar-title>
            
            <v-spacer></v-spacer>
            
            <v-menu>
                <template v-slot:activator="{ props }">
                    <v-btn variant="outlined" color="grey-darken-2" v-bind="props">
                        <span>{{ currentViewLabel }}</span>
                        <v-icon :icon="mdiMenuDown" end></v-icon>
                    </v-btn>
                </template>
                <v-list>
                    <v-list-item @click="viewType = 'month'">
                        <v-list-item-title>Month</v-list-item-title>
                    </v-list-item>
                </v-list>
            </v-menu>
        </v-toolbar>

        <!-- Calendar Days Header -->
        <div class="calendar-header-grid mt-4">
            <div class="text-center font-weight-bold text-caption text-grey-darken-1 py-2" v-for="day in weekDays" :key="day">
                {{ day }}
            </div>
        </div>

        <!-- Calendar Grid -->
        <div class="calendar-grid bg-grey-lighten-3">
            <!-- Empty offset days for the start of the month -->
            <div v-for="blank in blankDaysOffset" :key="'blank-' + blank" class="calendar-cell blank-cell bg-white"></div>
            
            <!-- Actual days of the month -->
            <div v-for="day in daysInMonth" :key="day.date" class="calendar-cell bg-white" :class="{ 'bg-primary-lighten-5': day.isToday }">
                <div class="d-flex justify-space-between align-center pa-2">
                    <span class="day-number font-weight-medium" :class="{ 'text-primary rounded-circle bg-primary-lighten-4 px-2': day.isToday }">
                        {{ day.dayNumber }}
                    </span>
                </div>
                
                <!-- Events -->
                <div class="events-container px-1 pb-1">
                    <div v-for="(event, idx) in day.events" :key="idx" 
                         class="event-chip text-truncate text-caption px-2 py-1 mb-1 rounded"
                         :style="{ backgroundColor: event.color, color: 'white' }"
                         @click="showEventDetails(event, $event)">
                        {{ event.name }}
                    </div>
                </div>
            </div>
        </div>

        <!-- Event Details Menu (Adapted from user snippet) -->
        <v-menu v-model="selectedOpen" :activator="selectedElement" :close-on-content-click="false" offset-y>
            <v-card min-width="350px" class="rounded-lg shadow-soft">
                <v-toolbar :color="selectedEvent?.color || 'primary'" dark class="text-white">
                    <v-btn icon variant="text"><v-icon :icon="mdiPencil"></v-icon></v-btn>
                    <v-toolbar-title class="text-white">{{ selectedEvent?.name }}</v-toolbar-title>
                    <v-spacer></v-spacer>
                    <v-btn icon variant="text"><v-icon :icon="mdiHeartOutline"></v-icon></v-btn>
                    <v-btn icon variant="text"><v-icon :icon="mdiDotsVertical"></v-icon></v-btn>
                </v-toolbar>
                <v-card-text class="pt-4">
                    <div class="d-flex align-center mb-2">
                        <v-icon :icon="mdiClockOutline" size="small" class="me-2 text-grey"></v-icon>
                        <span class="text-body-2">{{ selectedEvent?.start }} to {{ selectedEvent?.end }}</span>
                    </div>
                    <span v-html="selectedEvent?.details || 'No additional details available.'"></span>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn variant="text" color="grey-darken-1" @click="selectedOpen = false">Cancel</v-btn>
                </v-card-actions>
            </v-card>
        </v-menu>
    </v-card>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { 
    mdiChevronLeft, 
    mdiChevronRight, 
    mdiMenuDown, 
    mdiPencil, 
    mdiHeartOutline, 
    mdiDotsVertical,
    mdiClockOutline
} from '@mdi/js';
import { useI18n } from '@/locales/helpers.ts';

const { tt } = useI18n();

// --- STATE ---
const currentDate = ref(new Date());
const viewType = ref('month');
const currentViewLabel = computed(() => 'Month'); // Simplified for now since we built the Month grid

const selectedOpen = ref(false);
const selectedEvent = ref<any>(null);
const selectedElement = ref<HTMLElement | null>(null);

// Placeholder arrays mapped from the Vuetify 2 snippet
const colors = ['#0DAD83', '#FF6E40', '#3B82F6', '#8B5CF6', '#10B981', '#F59E0B'];
const names = ['Meeting', 'Holiday', 'PTO', 'Travel', 'Event', 'Birthday', 'Conference', 'Party'];

// --- GETTERS ---
const currentYear = computed(() => currentDate.value.getFullYear());
const currentMonthNumber = computed(() => currentDate.value.getMonth());
const currentMonthName = computed(() => {
    return currentDate.value.toLocaleString('default', { month: 'long' });
});

// Calculate the grid structure
const weekDays = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];

const blankDaysOffset = computed(() => {
    // Determine what day of the week the 1st of the month falls on
    const firstDayOfMonth = new Date(currentYear.value, currentMonthNumber.value, 1).getDay();
    return firstDayOfMonth; // Returns 0-6 (Sun-Sat)
});

const daysInMonth = computed(() => {
    const days = [];
    // Number of days in the current month
    const totalDays = new Date(currentYear.value, currentMonthNumber.value + 1, 0).getDate();
    
    const today = new Date();
    const isCurrentMonth = today.getMonth() === currentMonthNumber.value && today.getFullYear() === currentYear.value;

    for (let i = 1; i <= totalDays; i++) {
        // Generate mock events for demonstration (as the Vue 2 snippet did realistically)
        const dailyEvents = [];
        if (Math.random() > 0.6) { // 40% chance to have an event
            dailyEvents.push({
                name: names[Math.floor(Math.random() * names.length)],
                start: `${i}/${currentMonthNumber.value + 1} 10:00 AM`,
                end: `${i}/${currentMonthNumber.value + 1} 11:30 AM`,
                color: colors[Math.floor(Math.random() * colors.length)]
            });
        }
        if (Math.random() > 0.85) { // 15% chance to have a second event
            dailyEvents.push({
                name: names[Math.floor(Math.random() * names.length)],
                start: `${i}/${currentMonthNumber.value + 1} 2:00 PM`,
                end: `${i}/${currentMonthNumber.value + 1} 3:00 PM`,
                color: colors[Math.floor(Math.random() * colors.length)]
            });
        }

        days.push({
            date: new Date(currentYear.value, currentMonthNumber.value, i),
            dayNumber: i,
            isToday: isCurrentMonth && today.getDate() === i,
            events: dailyEvents
        });
    }
    return days;
});

// --- ACTIONS ---
function prevMonth() {
    currentDate.value = new Date(currentYear.value, currentMonthNumber.value - 1, 1);
}

function nextMonth() {
    currentDate.value = new Date(currentYear.value, currentMonthNumber.value + 1, 1);
}

function setToday() {
    currentDate.value = new Date();
}

function showEventDetails(event: any, e: MouseEvent) {
    e.stopPropagation();
    selectedEvent.value = event;
    selectedElement.value = e.currentTarget as HTMLElement;
    selectedOpen.value = true;
}

onMounted(() => {
    // Initial fetch/setup logic here
});
</script>

<style scoped>
.calendar-header-grid {
    display: grid;
    grid-template-columns: repeat(7, 1fr);
    border-bottom: 1px solid #e0e0e0;
}

.calendar-grid {
    display: grid;
    grid-template-columns: repeat(7, 1fr);
    gap: 1px; /* Creates the border lines between days */
}

.calendar-cell {
    min-height: 120px;
    display: flex;
    flex-direction: column;
    transition: background-color 0.2s ease;
}

.calendar-cell:hover:not(.blank-cell) {
    background-color: #f8fafc !important; /* Soft hover effect */
}

.events-container {
    flex-grow: 1;
    overflow-y: auto;
    max-height: 80px;
}

/* Custom scrollbar for events */
.events-container::-webkit-scrollbar {
    width: 4px;
}
.events-container::-webkit-scrollbar-thumb {
    background: #cbd5e1;
    border-radius: 4px;
}

.event-chip {
    cursor: pointer;
    user-select: none;
    font-size: 0.75rem;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    transition: filter 0.2s;
}

.event-chip:hover {
    filter: brightness(1.1);
}

.bg-primary-lighten-5 {
    background-color: rgba(13, 173, 131, 0.05) !important;
}
.bg-primary-lighten-4 {
    background-color: rgba(13, 173, 131, 0.15) !important;
}
</style>
