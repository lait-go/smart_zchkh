<script setup lang="ts">
// скрипт оставлен пустым — здесь можно добавить логику или переходы позже
</script>

<template>
  <main class="welcome-page">
    <div class="welcome-card container">
      <!-- Анимированная IoT-иконка -->
      <div class="iot-animation">
        <svg viewBox="0 0 64 64" class="iot-icon" aria-hidden="true">
          <circle cx="32" cy="8" r="6" class="node node-1" />
          <circle cx="12" cy="48" r="6" class="node node-2" />
          <circle cx="52" cy="48" r="6" class="node node-3" />
          <line x1="32" y1="14" x2="12" y2="42" class="connection" />
          <line x1="32" y1="14" x2="52" y2="42" class="connection" />
          <line x1="12" y1="42" x2="52" y2="42" class="connection" />
        </svg>
        <div class="pulse-effect"></div>
      </div>

      <h1 class="welcome-title">Добро пожаловать в IoT Dashboard</h1>
      <p class="welcome-subtitle">Управляйте вашими устройствами в едином интерфейсе</p>

      <div class="welcome-actions">
        <RouterLink to="/login" class="action-button primary">Начать работу</RouterLink>
        <RouterLink to="/about" class="action-button secondary">О системе</RouterLink>
      </div>
    </div>

    <div class="decorative-elements">
      <div class="decor-circle circle-1"></div>
      <div class="decor-circle circle-2"></div>
      <div class="decor-circle circle-3"></div>
    </div>
  </main>
</template>

<style scoped>
.welcome-page {
  position: relative;
  overflow: hidden;
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem;
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.2) 0%, var(--color-bg-light) 100%);
  animation: fadeInUp 0.8s ease;
}

.welcome-card {
  position: relative;
  z-index: 10;
  max-width: 480px;
  width: 100%;
  padding: 3.5rem;
  background: var(--color-text-light);
  backdrop-filter: blur(8px);
  border: 1px solid rgba(59, 130, 246, 0.2);
  border-radius: 1.75rem;
  box-shadow: var(--shadow-lg);
  text-align: center;
  animation: cardEntrance 0.8s cubic-bezier(0.22, 1, 0.36, 1);
}

.iot-animation {
  position: relative;
  width: 6rem;
  height: 6rem;
  margin: 0 auto 2rem;
}

.iot-icon {
  width: 100%;
  height: 100%;
  z-index: 2;
  position: relative;
}

.iot-icon .node {
  fill: var(--color-primary);
  stroke: var(--color-primary-dark);
  stroke-width: 1.5;
  transform-origin: center;
  animation: nodePulse 2s infinite ease-in-out;
}

.iot-icon .node-1 {
  animation-delay: 0s;
}
.iot-icon .node-2 {
  animation-delay: 0.3s;
}
.iot-icon .node-3 {
  animation-delay: 0.6s;
}

.iot-icon .connection {
  stroke: var(--color-primary-light);
  stroke-width: 1.5;
  stroke-dasharray: 100;
  stroke-dashoffset: 100;
  animation: drawConnections 1.2s forwards ease-in-out;
}

.pulse-effect {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border-radius: 50%;
  background: rgba(59, 130, 246, 0.1);
  animation: pulse 3s infinite ease-out;
  z-index: 1;
}

.welcome-title {
  font-size: 2rem;
  font-weight: 700;
  color: var(--color-primary-dark);
  margin-bottom: 1rem;
  position: relative;
}

.welcome-subtitle {
  font-size: 1.125rem;
  color: var(--color-text-dark);
  opacity: 0.9;
  margin-bottom: 2rem;
}

.welcome-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
}

.action-button {
  padding: 0.75rem 1.75rem;
  border-radius: 0.75rem;
  font-weight: 500;
  text-decoration: none;
  transition: var(--transition-default);
  display: inline-flex;
  align-items: center;
  justify-content: center;
}

.action-button.primary {
  background: var(--color-primary);
  color: var(--color-text-light);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
}
.action-button.primary:hover {
  background: var(--color-primary-dark);
  transform: translateY(-2px);
}

.action-button.secondary {
  background: transparent;
  border: 1px solid var(--color-primary);
  color: var(--color-primary);
}
.action-button.secondary:hover {
  background: rgba(59, 130, 246, 0.1);
  transform: translateY(-2px);
}

.decorative-elements {
  position: absolute;
  inset: 0;
  pointer-events: none;
  z-index: 0;
}
.decor-circle {
  position: absolute;
  border-radius: 50%;
  background: rgba(59, 130, 246, 0.05);
  filter: blur(40px);
}

.circle-1 {
  top: -50px;
  right: -50px;
  width: 200px;
  height: 200px;
  animation: float 12s infinite ease-in-out;
}
.circle-2 {
  bottom: -100px;
  left: -100px;
  width: 300px;
  height: 300px;
  animation: float 15s infinite ease-in-out reverse;
}
.circle-3 {
  top: 40%;
  left: 20%;
  width: 150px;
  height: 150px;
  animation: float 10s infinite ease-in-out 2s;
}

/* Animations */
@keyframes cardEntrance {
  from {
    opacity: 0;
    transform: translateY(20px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}
@keyframes nodePulse {
  0%,
  100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.1);
  }
}
@keyframes drawConnections {
  to {
    stroke-dashoffset: 0;
  }
}
@keyframes pulse {
  0% {
    transform: scale(0.8);
    opacity: 0.7;
  }
  70% {
    transform: scale(1.3);
    opacity: 0;
  }
  100% {
    transform: scale(0.8);
    opacity: 0;
  }
}
@keyframes float {
  0%,
  100% {
    transform: translate(0, 0);
  }
  25% {
    transform: translate(10px, 15px);
  }
  50% {
    transform: translate(-5px, 20px);
  }
  75% {
    transform: translate(15px, -10px);
  }
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

/* Responsive */
@media (max-width: 768px) {
  .welcome-card {
    padding: 2.5rem;
    max-width: 90%;
  }
  .welcome-title {
    font-size: 1.75rem;
  }
  .welcome-subtitle {
    font-size: 1rem;
  }
  .welcome-actions {
    flex-direction: column;
  }
  .action-button {
    width: 100%;
  }
}
@media (max-width: 480px) {
  .welcome-card {
    padding: 2rem 1.5rem;
  }
  .iot-animation {
    width: 5rem;
    height: 5rem;
    margin-bottom: 1.5rem;
  }
}
</style>
