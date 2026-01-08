<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/shared/store/auth';
import { registerUser } from '@/shared/api';

const router = useRouter();
const auth = useAuthStore();

const username = ref('');
const password = ref('');
const confirmPassword = ref('');
const email = ref('');
const firstName = ref('');
const lastName = ref('');
const errors = ref<string[]>([]);
const message = ref('');

async function handleRegister() {
  errors.value = [];
  message.value = '';

  if (!username.value.trim()) errors.value.push('Введите логин');
  if (!email.value.trim()) errors.value.push('Введите email');
  if (!firstName.value.trim()) errors.value.push('Введите имя');
  if (!lastName.value.trim()) errors.value.push('Введите фамилию');
  if (password.value.length < 6) errors.value.push('Пароль должен быть не менее 6 символов');
  if (password.value !== confirmPassword.value) errors.value.push('Пароли не совпадают');
  if (errors.value.length) return;

  try {
    await registerUser(
      username.value,
      password.value,
      email.value,
      firstName.value,
      lastName.value,
    );
    message.value = '✅ Регистрация прошла успешно';
    auth.login(username.value, 'fake-token', auth.userId);
    router.push('/dashboard');
  } catch (err: any) {
    message.value = err.message || 'Ошибка регистрации';
  }
}
</script>

<template>
  <div class="page-wrapper container">
    <div class="form-card">
      <svg viewBox="0 0 24 24" class="form-icon" aria-hidden="true">
        <path d="M12 12a5 5 0 1 0-5-5 5 5 0 0 0 5 5zm0 2c-4 0-7 3-7 5v3h14v-3c0-2-3-5-7-5z" />
      </svg>

      <h1 class="form-title">Регистрация</h1>

      <form @submit.prevent="handleRegister" novalidate>
        <div class="form-group">
          <label for="username">Логин</label>
          <input
            id="username"
            type="text"
            v-model="username"
            placeholder="Введите логин"
            class="form-input"
          />
        </div>

        <div class="form-group">
          <label for="email">Email</label>
          <input
            id="email"
            type="email"
            v-model="email"
            placeholder="Введите email"
            class="form-input"
          />
        </div>

        <div class="form-group">
          <label for="firstName">Имя</label>
          <input
            id="firstName"
            type="text"
            v-model="firstName"
            placeholder="Введите имя"
            class="form-input"
          />
        </div>

        <div class="form-group">
          <label for="lastName">Фамилия</label>
          <input
            id="lastName"
            type="text"
            v-model="lastName"
            placeholder="Введите фамилию"
            class="form-input"
          />
        </div>

        <div class="form-group">
          <label for="password">Пароль</label>
          <input
            id="password"
            type="password"
            v-model="password"
            placeholder="Введите пароль"
            class="form-input"
          />
        </div>

        <div class="form-group">
          <label for="confirm">Повторите пароль</label>
          <input
            id="confirm"
            type="password"
            v-model="confirmPassword"
            placeholder="Повторите пароль"
            class="form-input"
          />
        </div>

        <button type="submit" class="submit-button">📝 Зарегистрироваться</button>
      </form>

      <ul v-if="errors.length" class="errors">
        <li v-for="(err, i) in errors" :key="i">{{ err }}</li>
      </ul>

      <p v-if="message" :class="message.startsWith('✅') ? 'message-success' : 'message-error'">
        {{ message }}
      </p>
    </div>
  </div>
</template>

<style scoped>
.page-wrapper {
  position: relative;
  overflow: hidden;
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem 1rem;
  background: linear-gradient(145deg, rgba(37, 99, 235, 0.2) 0%, rgba(37, 99, 235, 0.6) 100%);
  animation: fadeInUp 0.8s ease-out;
}
.page-wrapper::before {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle at center, rgba(255, 255, 255, 0.1) 0%, transparent 70%);
  animation: rotate 25s linear infinite;
  z-index: 0;
}

.form-card {
  position: relative;
  z-index: 10;
  width: 100%;
  max-width: 420px;
  padding: 3rem 2rem;
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(8px);
  border: 1px solid rgba(255, 255, 255, 0.4);
  border-radius: 1.5rem;
  box-shadow:
    0 12px 24px -6px rgba(0, 0, 0, 0.2),
    inset 0 1px 2px rgba(255, 255, 255, 0.6);
  text-align: center;
  animation: cardEntrance 0.8s cubic-bezier(0.22, 1, 0.36, 1);
}

.form-icon {
  width: 3rem;
  height: 3rem;
  margin: 0 auto 1rem;
  stroke: var(--color-primary);
  stroke-width: 2;
  fill: none;
  transition: var(--transition-default);
}
.form-icon:hover {
  transform: scale(1.2);
}

.form-title {
  font-size: 1.75rem;
  font-weight: 700;
  color: var(--color-primary-dark);
  margin-bottom: 1.25rem;
  position: relative;
}
.form-title::after {
  content: '';
  position: absolute;
  bottom: -6px;
  left: 50%;
  transform: translateX(-50%);
  width: 45px;
  height: 3px;
  background: var(--color-primary);
  border-radius: 2px;
  box-shadow: 0 0 6px rgba(37, 99, 235, 0.5);
}

.form-group {
  margin-bottom: 1.5rem;
  text-align: left;
}

label {
  display: block;
  margin-bottom: 0.4rem;
  color: var(--color-text-dark);
  font-weight: 600;
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
  width: 100%;
  padding: 0.75rem;
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

.errors {
  margin-top: 1rem;
  list-style: none;
  color: var(--color-error);
  font-size: 0.875rem;
  text-align: left;
  padding-left: 1rem;
}

.message-success {
  margin-top: 1rem;
  color: var(--color-success);
  font-weight: 600;
}

.message-error {
  margin-top: 1rem;
  color: var(--color-error);
  font-weight: 600;
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

/* Responsive */
@media (max-width: 480px) {
  .form-card {
    padding: 2.5rem 1.5rem;
  }
  .form-title {
    font-size: 1.5rem;
  }
  .form-input {
    padding: 0.6rem 0.8rem;
    font-size: 0.95rem;
  }
  .submit-button {
    font-size: 0.95rem;
  }
}
</style>
