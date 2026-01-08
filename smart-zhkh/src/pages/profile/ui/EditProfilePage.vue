<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { fetchAccountById, updateAccount } from '@/shared/api/accounts';

const route = useRoute();
const router = useRouter();
const account = ref<any | null>(null);

onMounted(async () => {
  const id = Number(route.params.id);
  try {
    account.value = await fetchAccountById(id);
  } catch (e) {
    console.error('Ошибка загрузки счёта:', e);
  }
});

async function submitForm() {
  try {
    await updateAccount(account.value);
    router.push('/profile');
  } catch (e) {
    alert('❌ Ошибка при сохранении');
  }
}
</script>

<template>
  <div class="edit-account-container container">
    <h2 class="edit-title">
      <svg viewBox="0 0 24 24" class="edit-icon" aria-hidden="true">
        <path d="M4 21v-4.586l8.293-8.293 4.586 4.586L8.586 21H4z" />
        <path d="M20.707 7.293a1 1 0 0 0-1.414 0l-2 2    1.414 1.414 2-2a1 1 0 0 0 0-1.414z" />
      </svg>
      Редактировать счёт
    </h2>

    <form v-if="account" @submit.prevent="submitForm" class="account-form">
      <div class="form-group">
        <label for="full_name">ФИО</label>
        <input id="full_name" v-model="account.full_name" type="text" required class="form-input" />
      </div>

      <div class="form-group">
        <label for="address">Адрес</label>
        <input id="address" v-model="account.address" type="text" required class="form-input" />
      </div>

      <div class="form-group">
        <label for="area">Площадь (м²)</label>
        <input
          id="area"
          v-model.number="account.area"
          type="number"
          min="1"
          required
          class="form-input"
        />
      </div>

      <button type="submit" class="btn btn-primary submit-button">💾 Сохранить изменения</button>
    </form>

    <div v-else class="loading">Загрузка счёта...</div>
  </div>
</template>

<style scoped>
.edit-account-container {
  margin: 3rem auto;
  padding: 2.5rem 1.5rem;
  background: linear-gradient(145deg, var(--color-bg-light) 0%, var(--color-primary-light) 100%);
  border-radius: 1.5rem;
  box-shadow: var(--shadow-lg);
  animation: fadeIn 0.5s ease;
  color: var(--color-text-dark);
}

.edit-title {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-size: 1.75rem;
  color: var(--color-primary-dark);
  margin-bottom: 1.5rem;
}

.edit-icon {
  width: 2rem;
  height: 2rem;
  stroke: var(--color-primary);
  stroke-width: 2;
  fill: none;
  transition: var(--transition-default);
}
.edit-icon:hover {
  transform: scale(1.1);
}

.account-form {
  display: grid;
  gap: 1.25rem;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-group label {
  margin-bottom: 0.5rem;
  font-weight: 600;
  color: var(--color-text-dark);
}

.form-input {
  width: 100%;
  padding: 0.75rem 1rem;
  border: 1px solid var(--color-primary-light);
  border-radius: 0.5rem;
  font-size: 1rem;
  background: var(--color-bg-light);
  transition: var(--transition-default);
}
.form-input:focus {
  border-color: var(--color-primary);
  outline: none;
  background: var(--color-text-light);
}

.submit-button {
  justify-self: end;
  padding: 0.75rem 1.5rem;
}

.btn-primary {
  background-color: var(--color-primary);
  color: var(--color-text-light);
  transition: var(--transition-default);
}
.btn-primary:hover {
  background-color: var(--color-primary-dark);
}

.loading {
  text-align: center;
  color: var(--color-text-dark);
  font-size: 1rem;
  padding: 2rem 0;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(12px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Адаптивность */
@media (max-width: 768px) {
  .edit-account-container {
    padding: 2rem 1rem;
  }
  .edit-title {
    font-size: 1.5rem;
  }
  .submit-button {
    justify-self: stretch;
  }
}

@media (max-width: 480px) {
  .form-input {
    font-size: 0.95rem;
    padding: 0.6rem 0.8rem;
  }
  .edit-title {
    font-size: 1.25rem;
  }
}
</style>
