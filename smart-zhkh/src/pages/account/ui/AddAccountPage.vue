<script setup lang="ts">
import { reactive } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/shared/store/auth';
import { createAccount } from '@/shared/api/accounts';

const auth = useAuthStore();
const router = useRouter();

const form = reactive({
  account_number: '',
  full_name: '',
  address: '',
  area: 0,
});

async function onSubmit() {
  try {
    await createAccount({
      user_id: auth.userId,
      account_number: form.account_number,
      full_name: form.full_name,
      address: form.address,
      area: form.area,
    });
    alert('✅ Счёт успешно создан');
    router.push('/profile');
  } catch (e: any) {
    alert('❌ Ошибка: ' + (e.message || 'Не удалось добавить счёт'));
  }
}
</script>

<template>
  <div class="add-account-container container">
    <h2 class="form-title">
      <svg viewBox="0 0 24 24" class="form-icon" aria-hidden="true">
        <path d="M12 12a5 5 0 1 0-5-5 5 5 0 0 0 5 5zm0 2c-4 0-7 3-7 5v2h14v-2c0-2-3-5-7-5z" />
      </svg>
      Добавить счёт
    </h2>

    <form @submit.prevent="onSubmit" class="account-form">
      <div class="form-group">
        <label for="account_number">Номер счёта</label>
        <input
          id="account_number"
          v-model="form.account_number"
          type="text"
          required
          class="form-input"
          placeholder="Введите номер счёта"
        />
      </div>

      <div class="form-group">
        <label for="full_name">ФИО</label>
        <input
          id="full_name"
          v-model="form.full_name"
          type="text"
          required
          class="form-input"
          placeholder="Введите ФИО"
        />
      </div>

      <div class="form-group">
        <label for="address">Адрес</label>
        <input
          id="address"
          v-model="form.address"
          type="text"
          required
          class="form-input"
          placeholder="Введите адрес"
        />
      </div>

      <div class="form-group">
        <label for="area">Площадь (м²)</label>
        <input
          id="area"
          v-model.number="form.area"
          type="number"
          min="1"
          required
          class="form-input"
          placeholder="Введите площадь"
        />
      </div>

      <button type="submit" class="submit-button">Сохранить</button>
    </form>
  </div>
</template>

<style scoped>
.add-account-container {
  position: relative;
  overflow: hidden;
  margin: 3rem auto;
  padding: 3rem 2rem;
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.4);
  border-radius: 1.75rem;
  box-shadow:
    0 12px 30px -8px rgba(0, 0, 0, 0.2),
    inset 0 2px 4px rgba(255, 255, 255, 0.6);
  animation: cardEntrance 0.8s cubic-bezier(0.22, 1, 0.36, 1);
  color: var(--color-text-dark);
}
.add-account-container::before {
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

.form-title {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-size: 1.75rem;
  font-weight: 800;
  color: var(--color-primary-dark);
  margin-bottom: 2rem;
  position: relative;
}
.form-title::after {
  content: '';
  position: absolute;
  bottom: -6px;
  left: 0;
  width: 50px;
  height: 4px;
  background: var(--color-primary-light);
  border-radius: 2px;
  box-shadow: 0 0 8px var(--color-primary-light);
}

.form-icon {
  width: 2rem;
  height: 2rem;
  stroke: var(--color-primary);
  stroke-width: 2;
  fill: none;
  transition: transform 0.3s;
}
.form-icon:hover {
  transform: scale(1.1);
}

.account-form {
  display: grid;
  gap: 1.5rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 600;
  color: var(--color-text-dark);
}

.form-input {
  width: 100%;
  padding: 0.75rem 1rem;
  border: 1px solid var(--color-primary-light);
  border-radius: 0.75rem;
  font-size: 1rem;
  background: var(--color-text-light);
  transition: var(--transition-default);
}
.form-input:focus {
  border-color: var(--color-primary);
  outline: none;
  background: var(--color-bg-light);
}

.submit-button {
  justify-self: end;
  padding: 0.75rem 1.5rem;
  background: var(--color-primary);
  color: var(--color-text-light);
  border: none;
  border-radius: 0.75rem;
  font-size: 1rem;
  font-weight: 600;
  transition: var(--transition-default);
}
.submit-button:hover {
  background: var(--color-primary-dark);
  transform: translateY(-2px);
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
  .add-account-container {
    padding: 2.5rem 1.5rem;
  }
  .form-title {
    font-size: 1.5rem;
  }
  .submit-button {
    width: 100%;
  }
}
@media (max-width: 480px) {
  .form-input {
    padding: 0.6rem 0.8rem;
    font-size: 0.95rem;
  }
  .form-title {
    font-size: 1.25rem;
  }
  .submit-button {
    font-size: 0.95rem;
  }
}
</style>
