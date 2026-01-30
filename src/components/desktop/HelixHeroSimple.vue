<template>
  <div class="helix-hero-container">
    
    <div class="animated-background">
      <div class="helix-ring" v-for="i in 18" :key="i" :style="getRingStyle(i)"></div>
    </div>

    
    <div class="glow-sphere"></div>
    <div class="gradient-overlay gradient-top"></div>
    <div class="gradient-overlay gradient-bottom"></div>

    
    <div class="helix-content animate-fade-in-up">
      <h1 class="helix-title">{{ title }}</h1>
      <p class="helix-description">{{ description }}</p>
      
      <div class="d-flex align-center mt-6">
        <div class="status-dot"></div>
        <span class="text-caption text-white opacity-60 font-weight-medium uppercase tracking-wider ms-2">System Active</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  title?: string;
  description?: string;
}

const props = withDefaults(defineProps<Props>(), {
  title: 'Welcome to Split Chill AI',
  description: 'Experience seamless expense tracking with cutting-edge technology. Manage your finances with intelligence and style.'
});

const getRingStyle = (index: number) => {
  const delay = index * 0.15;
  const scale = 0.4 + (index * 0.08);
  const rotation = index * 20;
  
  return {
    animationDelay: `${delay}s`,
    transform: `scale(${scale}) rotate(${rotation}deg)`,
    opacity: 0.05 + (index * 0.015)
  };
};
</script>

<style scoped>
.helix-hero-container {
  position: relative;
  width: 100%;
  height: 100%;
  overflow: hidden;
  background: radial-gradient(circle at 20% 30%, hsla(221, 83%, 93%, 1), transparent),
              radial-gradient(circle at 80% 70%, hsla(142, 71%, 95%, 1), transparent),
              #f8fafc;
}

.animated-background {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 100%;
  height: 100%;
  z-index: 1;
}

.helix-ring {
  position: absolute;
  top: 50%;
  left: 50%;
  width: 240px;
  height: 240px;
  margin: -120px 0 0 -120px;
  border: 4px solid rgba(var(--ebk-primary-h, 221), var(--ebk-primary-s, 83%), var(--ebk-primary-l, 53%), 0.1);
  border-radius: 42% 58% 70% 30% / 45% 45% 55% 55%;
  animation: rotate 25s linear infinite, pulse 4s ease-in-out infinite, morph 8s ease-in-out infinite;
  box-shadow: 0 0 30px hsla(221, 83%, 53%, 0.1);
}

.glow-sphere {
  position: absolute;
  top: 40%;
  left: 50%;
  width: 400px;
  height: 400px;
  background: radial-gradient(circle, hsla(221, 83%, 53%, 0.15), transparent 70%);
  filter: blur(80px);
  z-index: 2;
  pointer-events: none;
}

@keyframes rotate {
  from { transform: rotate(0deg) scale(var(--scale, 1)); }
  to { transform: rotate(360deg) scale(var(--scale, 1)); }
}

@keyframes morph {
  0%, 100% { border-radius: 42% 58% 70% 30% / 45% 45% 55% 55%; }
  50% { border-radius: 70% 30% 46% 54% / 30% 29% 71% 70%; }
}

@keyframes pulse {
  0%, 100% {
    border-color: hsla(221, 83%, 53%, 0.1);
    box-shadow: 0 0 20px hsla(142, 71%, 45%, 0.1);
  }
  50% {
    border-color: hsla(142, 71%, 45%, 0.4);
    box-shadow: 0 0 50px hsla(221, 83%, 53%, 0.2);
  }
}

.gradient-overlay {
  position: absolute;
  inset: 0;
  z-index: 3;
  pointer-events: none;
}

.gradient-top {
  height: 40%;
  background: linear-gradient(to bottom, rgba(248, 250, 252, 0.6), transparent);
}

.gradient-bottom {
  top: auto;
  bottom: 0;
  height: 60%;
  background: linear-gradient(to top, rgba(248, 250, 252, 0.8), transparent);
}

.helix-content {
  position: absolute;
  bottom: 3rem;
  left: 3.5rem;
  z-index: 10;
  max-width: 32rem;
  background: rgba(15, 23, 42, 0.8);
  padding: 2.5rem;
  border-radius: 28px;
  backdrop-filter: blur(24px);
  -webkit-backdrop-filter: blur(24px);
  border: 1px solid rgba(255, 255, 255, 0.08);
  box-shadow: 0 24px 80px rgba(0, 0, 0, 0.25);
}

.helix-title {
  font-family: 'Outfit', sans-serif;
  font-size: 3rem;
  font-weight: 800;
  letter-spacing: -0.03em;
  margin-bottom: 1rem;
  color: #FFFFFF;
  line-height: 1.1;
  background: linear-gradient(to bottom, #fff, rgba(255,255,255,0.7));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.helix-description {
  color: rgba(255, 255, 255, 0.7);
  font-family: 'Inter', sans-serif;
  font-size: 1.125rem;
  line-height: 1.6;
  font-weight: 400;
}

.status-dot {
  width: 8px;
  height: 8px;
  background-color: hsla(142, 71%, 45%, 1);
  border-radius: 50%;
  box-shadow: 0 0 10px hsla(142, 71%, 45%, 0.6);
  animation: blink 2s infinite;
}

@keyframes blink {
  0%, 100% { opacity: 1; transform: scale(1); }
  50% { opacity: 0.4; transform: scale(0.8); }
}

.animate-fade-in-up {
  animation: fadeInUp 1.2s cubic-bezier(0.16, 1, 0.3, 1);
}

@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(30px); }
  to { opacity: 1; transform: translateY(0); }
}

.uppercase {
  text-transform: uppercase;
}

.tracking-wider {
  letter-spacing: 0.1em;
}
</style>
