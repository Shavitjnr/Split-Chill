<template>
  <TresGroup ref="groupRef" :scale="1" :position="[5, 0, 0]" :rotation="[0, 0, 0]">
    <TresMesh
      v-for="ring in rings"
      :key="ring.id"
      :geometry="ringGeometry"
      :position="[0, ring.y, 0]"
      :rotation="[0, Math.PI / 2 + ring.rotation, 0]"
    >
      <TresMeshStandardMaterial
        :color="'#3B82F6'"
        :metalness="0.8"
        :roughness="0.2"
        :emissive="'#3B82F6'"
        :emissive-intensity="0.2"
      />
    </TresMesh>
  </TresGroup>
</template>

<script setup lang="ts">
import { useLoop } from '@tresjs/core';
import { ref, computed } from 'vue';
import * as THREE from 'three';

interface Props {
  levelsUp?: number;
  levelsDown?: number;
  stepY?: number;
  rotationStep?: number;
}

const props = withDefaults(defineProps<Props>(), {
  levelsUp: 10,
  levelsDown: 10,
  stepY: 0.85,
  rotationStep: Math.PI / 16,
});

const groupRef = ref();


const { onBeforeRender } = useLoop();
onBeforeRender(() => {
  if (groupRef.value) {
    groupRef.value.rotation.y += 0.005;
  }
});


const ringGeometry = computed(() => {
  const shape = new THREE.Shape();
  const radius = 0.35;
  shape.absarc(0, 0, radius, 0, Math.PI * 2, false);

  const depth = 10;
  const extrudeSettings: THREE.ExtrudeGeometryOptions = {
    depth,
    bevelEnabled: true,
    bevelThickness: 0.05,
    bevelSize: 0.05,
    bevelSegments: 4,
    curveSegments: 64,
  };

  const geometry = new THREE.ExtrudeGeometry(shape, extrudeSettings);
  geometry.translate(0, 0, -depth / 2);

  return geometry;
});


const rings = computed(() => {
  const elements = [];
  for (let i = -props.levelsDown; i <= props.levelsUp; i++) {
    elements.push({
      id: `helix-ring-${i}`,
      y: i * props.stepY,
      rotation: i * props.rotationStep,
    });
  }
  return elements;
});
</script>
