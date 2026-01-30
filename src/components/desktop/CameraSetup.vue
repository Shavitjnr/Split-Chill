<script setup lang="ts">
import { useTresContext } from '@tresjs/core';
import { onMounted, watch } from 'vue';
import type { OrthographicCamera } from 'three';

interface Props {
  position?: [number, number, number];
  zoom?: number;
}

const props = withDefaults(defineProps<Props>(), {
  position: () => [0, 0, 7],
  zoom: 70,
});

const context = useTresContext() as any;

onMounted(() => {
  const cam = context?.camera?.activeCamera?.value as OrthographicCamera;
  if (cam) {
    cam.position.set(...props.position);
    cam.zoom = props.zoom;
    cam.updateProjectionMatrix();
  }
});

watch(() => props.position, (newPos) => {
  const cam = context?.camera?.activeCamera?.value as OrthographicCamera;
  if (cam) {
    cam.position.set(...newPos);
    cam.updateProjectionMatrix();
  }
});

watch(() => props.zoom, (newZoom) => {
  const cam = context?.camera?.activeCamera?.value as OrthographicCamera;
  if (cam) {
    cam.zoom = newZoom;
    cam.updateProjectionMatrix();
  }
});
</script>

<template>
  
</template>
