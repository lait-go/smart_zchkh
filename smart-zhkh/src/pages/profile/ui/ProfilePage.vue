<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useAuthStore } from '@/shared/store/auth';
import { fetchAccounts } from '@/shared/api/accounts';

const auth = useAuthStore();
const accounts = ref<any[]>([]);
const isLoading = ref(true);

onMounted(async () => {
  try {
    accounts.value = await fetchAccounts();
  } catch (e) {
    console.error('Ошибка загрузки счетов:', e);
  } finally {
    isLoading.value = false;
  }
});

const filteredAccounts = computed(() =>
  accounts.value.filter((acc) => String(acc.user_id) === String(auth.userId)),
);

function formatDate(date: string) {
  return new Date(date).toLocaleDateString('ru-RU');
}
</script>

<template>
  <div class="profile-container container">
    <h2 class="profile-title">
      <svg viewBox="0 0 24 24" class="profile-icon" aria-hidden="true">
        <path d="M12 12a5 5 0 1 0-5-5 5 5 0 0 0 5 5zm0 2c-4 0-7 3-7 5v3h14v-3c0-2-3-5-7-5z" />
      </svg>
      Профиль
    </h2>

    <div class="user-info">
      <p><strong>Пользователь:</strong> {{ auth.username }}</p>
    </div>
    <h3 class="section-title">
      Ваши счета
      <RouterLink to="/profile/history" class="btn btn-primary" style="margin-left: 1rem">
        📋 История заявок
      </RouterLink>
    </h3>

    <div v-if="isLoading" class="loading">⏳ Загрузка...</div>

    <div v-else-if="filteredAccounts.length === 0" class="empty-state">
      <p>У вас пока нет заполненного профиля.</p>
      <RouterLink to="/accounts/add" class="btn btn-primary mt-4">
        📝 Заполнить профиль
      </RouterLink>
    </div>

    <div v-else class="account-list">
      <div v-for="acc in filteredAccounts" :key="acc.id" class="account-card">
        <div class="card-topbar"></div>
        <svg viewBox="0 0 24 24" class="card-icon" aria-hidden="true">
          <path d="M3 6h18v12H3z" />
          <path d="M3 10h18" />
        </svg>
        <p><strong>Номер счёта:</strong> {{ acc.account_num }}</p>
        <p><strong>ФИО:</strong> {{ acc.full_name }}</p>
        <p><strong>Адрес:</strong> {{ acc.address }}</p>
        <p><strong>Площадь:</strong> {{ acc.area }} м²</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.profile-container {
  position: relative;
  overflow: hidden;
  margin: 2rem auto;
  padding: 3rem 2rem;
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.4);
  border-radius: 1.75rem;
  box-shadow:
    0 12px 30px -8px rgba(0, 0, 0, 0.2),
    inset 0 2px 4px rgba(255, 255, 255, 0.6);
  color: var(--color-text-dark);
  animation: cardEntrance 0.8s cubic-bezier(0.22, 1, 0.36, 1);
}
.profile-container::before {
  content: '';
  position: absolute;
  top: -40%;
  left: -40%;
  width: 180%;
  height: 180%;
  background: radial-gradient(circle at center, rgba(37, 99, 235, 0.1) 0%, transparent 70%);
  animation: rotate 30s linear infinite;
  z-index: -1;
}

.profile-title {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-size: 2rem;
  font-weight: 800;
  margin-bottom: 1rem;
  position: relative;
}
.profile-title::after {
  content: '';
  position: absolute;
  bottom: -6px;
  left: 0;
  width: 60px;
  height: 4px;
  background: var(--color-primary-light);
  border-radius: 2px;
  box-shadow: 0 0 8px var(--color-primary-light);
}
.profile-icon {
  width: 2.5rem;
  height: 2.5rem;
  fill: var(--color-primary);
  transition: transform 0.3s;
}
.profile-icon:hover {
  transform: scale(1.1);
}

.user-info {
  background: var(--color-bg-light);
  padding: 1rem;
  border-radius: 0.75rem;
  border-left: 4px solid var(--color-primary);
  margin-bottom: 2rem;
  color: var(--color-text-dark);
}

.section-title {
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 1rem;
  color: var(--color-text-dark);
}

.loading,
.empty-state {
  text-align: center;
  padding: 2rem 0;
  color: var(--color-text-dark);
}

.empty-state {
  border: 2px dashed var(--color-warning);
  border-radius: 1rem;
  background: var(--color-bg-light);
  color: var(--color-warning);
}

.btn {
  display: inline-flex;
  align-items: center;
  padding: 0.5rem 1rem;
  border-radius: 0.75rem;
  font-weight: 500;
}
.btn-primary {
  background: var(--color-primary);
  color: var(--color-text-light);
  transition: var(--transition-default);
}
.btn-primary:hover {
  background: var(--color-primary-dark);
}

.mt-4 {
  margin-top: 1rem;
}

.account-list {
  display: grid;
  gap: 2rem;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
}

.account-card {
  position: relative;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(8px);
  border-radius: 1.25rem;
  box-shadow:
    0 6px 16px -4px rgba(0, 0, 0, 0.1),
    inset 0 1px 3px rgba(255, 255, 255, 0.5);
  padding: 2rem 1.5rem 1.5rem;
  text-align: center;
  transition:
    transform 0.4s,
    box-shadow 0.4s;
}
.account-card:hover {
  transform: translateY(-6px) scale(1.02);
  box-shadow:
    0 12px 24px -6px rgba(0, 0, 0, 0.15),
    inset 0 1px 4px rgba(255, 255, 255, 0.6);
}

.card-topbar {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, var(--color-primary-light), var(--color-primary-dark));
}

.card-icon {
  width: 2rem;
  height: 2rem;
  margin-bottom: 1rem;
  stroke: var(--color-primary);
  stroke-width: 2;
  fill: none;
  transition:
    transform 0.3s,
    filter 0.3s;
  filter: drop-shadow(0 2px 4px rgba(37, 99, 235, 0.2));
}
.account-card:hover .card-icon {
  transform: scale(1.2) translateY(-4px);
  filter: drop-shadow(0 4px 8px rgba(37, 99, 235, 0.3));
}

.account-card p {
  margin: 0.25rem 0;
  color: var(--color-text-dark);
  font-size: 0.95rem;
}

@keyframes cardEntrance {
  from {
    opacity: 0;
    transform: translateY(30px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
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
  .profile-container {
    padding: 2rem 1rem;
  }
  .profile-title {
    font-size: 1.5rem;
  }
  .account-list {
    gap: 1.5rem;
  }
}
@media (max-width: 480px) {
  .profile-container {
    padding: 1.5rem;
  }
  .profile-title {
    font-size: 1.25rem;
  }
  .account-card {
    padding: 1.5rem 1rem 1rem;
  }
}
</style>
