<template>
  <div class="helix-hero-container">
    
    <TresCanvas
      v-bind="gl"
      class="helix-canvas"
      :orthographic="true"
    >
      
      <CameraSetup :position="[0, 0, 7]" :zoom="70" />

      
      <TresAmbientLight :intensity="0.5" />
      <TresDirectionalLight
        :position="[10, 10, 5]"
        :intensity="1.5"
        color="#ffffff"
      />

      
      <HelixRings />
    </TresCanvas>

    
    <div class="gradient-overlay gradient-top"></div>
    <div class="gradient-overlay gradient-bottom"></div>

    
    <div class="helix-content">
      <h1 class="helix-title">{{ title }}</h1>
      <p class="helix-description">{{ description }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { TresCanvas } from '@tresjs/core';
import { ref } from 'vue';
import * as THREE from 'three';
import HelixRings from './HelixRings.vue';
import CameraSetup from './CameraSetup.vue';

interface Props {
  title?: string;
  description?: string;
}

const props = withDefaults(defineProps<Props>(), {
  title: 'Welcome to Split Chill AI',
  description: 'Experience seamless expense tracking with cutting-edge technology. A silent rhythm spirals endlessly through empty space — light refracts, forms bend, and geometry hums in quiet harmony.'
});

const gl = ref({
  antialias: true,
  alpha: true,
  toneMapping: THREE.ACESFilmicToneMapping,
  toneMappingExposure: 1.2,
  outputColorSpace: THREE.SRGBColorSpace,
});
</script>

<style scoped>
.helix-hero-container {
  position: relative;
  width: 100%;
  height: 100%;
  overflow: hidden;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
}

.helix-canvas {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 1;
}

.gradient-overlay {
  position: absolute;
  left: 0;
  right: 0;
  z-index: 2;
  pointer-events: none;
}

.gradient-top {
  top: 0;
  height: 33%;
  background: linear-gradient(to bottom, rgba(255, 255, 255, 0.3), transparent);
  -webkit-backdrop-filter: blur(20px);
  backdrop-filter: blur(20px);
}

.gradient-bottom {
  bottom: 0;
  height: 50%;
  background: linear-gradient(to top, rgba(255, 255, 255, 0.4), transparent);
  -webkit-backdrop-filter: blur(30px);
  backdrop-filter: blur(30px);
}

.helix-content {
  position: absolute;
  bottom: 2.5rem;
  left: 2.5rem;
  z-index: 10;
  max-width: 28rem;
  animation: fadeInUp 1s ease-out;
}

@media (max-width: 768px) {
  .helix-content {
    bottom: 1rem;
    left: 1rem;
    max-width: calc(100% - 2rem);
  }
}

.helix-title {
  font-size: 2rem;
  font-weight: 300;
  letter-spacing: -0.02em;
  margin-bottom: 0.75rem;
  color: #1a202c;
  line-height: 1.2;
}

.helix-description {
  color: #4a5568;
  font-size: 0.875rem;
  line-height: 1.6;
  font-weight: 300;
  letter-spacing: -0.01em;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
