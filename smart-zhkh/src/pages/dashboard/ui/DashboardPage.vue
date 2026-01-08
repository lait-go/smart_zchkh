<template>
  <div class="dashboard-container container">
    <h2 class="dashboard-title">👋 Привет, {{ auth.username }}</h2>
    <p class="dashboard-subtitle">Добро пожаловать в личный кабинет</p>

    <div class="card-grid">
      <RouterLink to="/charges" class="info-card">
        <svg viewBox="0 0 24 24" class="card-icon" aria-hidden="true">
          <path
            d="M20 7H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1zm-1 2v2H5V9h14zm0 4v2H5v-2h14z"
          />
        </svg>
        <span>Мои начисления</span>
      </RouterLink>

      <div class="info-card disabled">
        <svg viewBox="0 0 24 24" class="card-icon" aria-hidden="true">
          <path d="M12 2a5 5 0 0 0-5 5v4H5l7 7 7-7h-2V7a5 5 0 0 0-5-5z" />
        </svg>
        <span>Устройства (в разработке)</span>
      </div>

      <RouterLink to="/profile" class="info-card">
        <svg viewBox="0 0 24 24" class="card-icon" aria-hidden="true">
          <path d="M12 12a5 5 0 1 0-5-5 5 5 0 0 0 5 5zm0 2c-4 0-7 2-7 4v2h14v-2c0-2-3-4-7-4z" />
        </svg>
        <span>Профиль</span>
      </RouterLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/shared/store/auth';

const router = useRouter();
const auth = useAuthStore();

onMounted(() => {
  if (!auth.isLoggedIn) {
    router.push('/login');
  }
});
</script>

<style scoped>
.dashboard-container {
  background: linear-gradient(145deg, rgba(37, 99, 235, 0.95) 0%, rgba(29, 78, 216, 0.95) 100%);
  backdrop-filter: blur(8px);
  -webkit-backdrop-filter: blur(8px);
  color: var(--color-text-light);
  max-width: 1000px;
  margin: 2rem auto;
  padding: 3rem 2rem;
  border-radius: 24px;
  box-shadow:
    0 10px 25px -5px rgba(0, 0, 0, 0.2),
    0 8px 10px -6px rgba(0, 0, 0, 0.1),
    inset 0 1px 2px rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.15);
  animation: fadeInUp 0.6s cubic-bezier(0.22, 1, 0.36, 1);
  position: relative;
  overflow: hidden;
}

.dashboard-container::before {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle at center, rgba(255, 255, 255, 0.1) 0%, transparent 70%);
  animation: rotate 20s linear infinite;
  z-index: -1;
}

.dashboard-title {
  font-size: 2rem;
  font-weight: 800;
  margin-bottom: 0.75rem;
  letter-spacing: -0.5px;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  position: relative;
  display: inline-block;
}

.dashboard-title::after {
  content: '';
  position: absolute;
  bottom: -8px;
  left: 0;
  width: 50px;
  height: 3px;
  background: var(--color-text-light);
  border-radius: 3px;
  box-shadow: 0 0 8px rgba(255, 255, 255, 0.6);
}

.dashboard-subtitle {
  font-size: 1.1rem;
  margin-bottom: 2.5rem;
  opacity: 0.9;
  max-width: 600px;
  line-height: 1.6;
}

.card-grid {
  display: grid;
  gap: 2rem;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
}

.info-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 2.25rem 1.5rem;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(8px);
  -webkit-backdrop-filter: blur(8px);
  border-radius: 16px;
  box-shadow:
    0 6px 12px -2px rgba(0, 0, 0, 0.1),
    0 4px 6px -1px rgba(0, 0, 0, 0.05),
    inset 0 1px 1px rgba(255, 255, 255, 0.5);
  color: var(--color-text-dark);
  text-decoration: none;
  transition: all 0.4s cubic-bezier(0.22, 1, 0.36, 1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  text-align: center;
  min-height: 180px;
  position: relative;
  overflow: hidden;
}

.info-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, var(--color-primary-light), var(--color-primary-dark));
}

.info-card:hover {
  transform: translateY(-6px) scale(1.02);
  box-shadow:
    0 12px 20px -5px rgba(0, 0, 0, 0.15),
    0 8px 10px -4px rgba(0, 0, 0, 0.1);
  background: rgba(255, 255, 255, 0.95);
  border-color: rgba(255, 255, 255, 0.5);
}

.info-card span {
  font-weight: 600;
  font-size: 1.1rem;
  margin-top: 0.5rem;
  transition: color 0.3s ease;
}

.info-card:hover span {
  color: var(--color-primary-dark);
}

.info-card.disabled {
  opacity: 0.7;
  cursor: not-allowed;
  background: rgba(255, 255, 255, 0.8);
}

.info-card.disabled:hover {
  transform: none;
  box-shadow:
    0 6px 12px -2px rgba(0, 0, 0, 0.1),
    0 4px 6px -1px rgba(0, 0, 0, 0.05);
}

.card-icon {
  width: 2.75rem;
  height: 2.75rem;
  margin-bottom: 1rem;
  fill: var(--color-primary);
  transition: all 0.4s cubic-bezier(0.22, 1, 0.36, 1);
  filter: drop-shadow(0 2px 4px rgba(37, 99, 235, 0.2));
}

.info-card:hover .card-icon {
  transform: scale(1.15) translateY(-5px);
  filter: drop-shadow(0 4px 8px rgba(37, 99, 235, 0.3));
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

@keyframes rotate {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

@media (max-width: 768px) {
  .dashboard-container {
    margin: 1.5rem;
    padding: 2.25rem 1.5rem;
    border-radius: 20px;
  }

  .dashboard-title {
    font-size: 1.75rem;
  }

  .dashboard-subtitle {
    font-size: 1rem;
    margin-bottom: 2rem;
  }

  .card-grid {
    gap: 1.5rem;
  }

  .info-card {
    padding: 1.75rem 1.25rem;
    min-height: 160px;
  }
}

@media (max-width: 480px) {
  .dashboard-container {
    margin: 1rem;
    padding: 2rem 1.25rem;
    border-radius: 18px;
  }

  .dashboard-title {
    font-size: 1.5rem;
  }

  .card-grid {
    grid-template-columns: 1fr;
    gap: 1.25rem;
  }

  .info-card {
    padding: 1.5rem;
    min-height: auto;
  }

  .card-icon {
    width: 2.25rem;
    height: 2.25rem;
    margin-bottom: 0.75rem;
  }
}
</style>
