<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/shared/store/auth';
import { loginUser } from '@/shared/api';

const username = ref('');
const password = ref('');
const errors = ref<string[]>([]);
const message = ref('');

const auth = useAuthStore();
const router = useRouter();
const mode = ref<'user' | 'admin' | null>(null);

async function handleLogin() {
  errors.value = [];
  message.value = '';

  if (!username.value.trim()) errors.value.push('Введите логин');
  if (!password.value.trim()) errors.value.push('Введите пароль');
  else if (password.value.length < 6) errors.value.push('Пароль должен быть не менее 6 символов');

  if (errors.value.length) return;

  try {
    const res = await loginUser(username.value, password.value);
    auth.login(res.username, res.token || 'mock-token', res.userId);
    router.push('/dashboard');
  } catch (err: any) {
    message.value = '❌ ' + (err.message || 'Ошибка входа');
  }
}

function loginAsAdmin() {
  auth.login('Администратор', 'admin-token', 'admin-001');
  router.push('/admin/charges');
}
</script>

<template>
  <div class="page-wrapper container">
    <div class="form-card">
      <svg viewBox="0 0 24 24" class="form-icon" aria-hidden="true">
        <rect x="5" y="11" width="14" height="10" rx="2" ry="2" />
        <path d="M12 11V7a4 4 0 0 0-8 0v4" />
        <path d="M12 11V7a4 4 0 0 1 8 0v4" />
      </svg>

      <h1 class="form-title">Вход в систему</h1>

      <div class="button-group">
        <button class="button user-button" @click="mode = 'user'">🔒 Войти как пользователь</button>
        <button class="button admin-button" @click="loginAsAdmin">
          🛠 Войти как администратор
        </button>
      </div>

      <form v-if="mode === 'user'" @submit.prevent="handleLogin" novalidate>
        <div class="form-group">
          <label for="username">Логин</label>
          <input
            v-model="username"
            id="username"
            type="text"
            placeholder="Введите логин"
            class="form-input"
          />
        </div>

        <div class="form-group">
          <label for="password">Пароль</label>
          <input
            v-model="password"
            id="password"
            type="password"
            placeholder="Введите пароль"
            class="form-input"
          />
        </div>

        <button type="submit" class="button submit-button">🔒 Войти</button>
      </form>

      <ul v-if="errors.length" class="errors">
        <li v-for="(err, i) in errors" :key="i">{{ err }}</li>
      </ul>

      <p v-if="message" class="message">{{ message }}</p>
    </div>
  </div>
</template>

<style scoped>
.page-wrapper {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem 1rem;
  background: linear-gradient(145deg, rgba(37, 99, 235, 0.95), rgba(29, 78, 216, 0.95));
}

.form-card {
  max-width: 400px;
  width: 100%;
  padding: 3rem 2rem;
  background: rgba(255, 255, 255, 0.97);
  border-radius: 1.5rem;
  box-shadow: 0 10px 20px -5px rgba(0, 0, 0, 0.2);
  text-align: center;
}

.form-icon {
  width: 3rem;
  height: 3rem;
  margin: 0 auto 1rem;
  stroke: #2563eb;
  stroke-width: 2;
  fill: none;
}

.form-title {
  font-size: 1.75rem;
  font-weight: 800;
  margin-bottom: 1.5rem;
}

.button-group {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  margin-bottom: 2rem;
}

.button {
  padding: 0.75rem;
  border: none;
  border-radius: 0.75rem;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.user-button {
  background: #3b82f6;
  color: white;
}
.user-button:hover {
  background: #2563eb;
}

.admin-button {
  background: #111827;
  color: white;
}
.admin-button:hover {
  background: #1f2937;
}

.submit-button {
  background: #16a34a;
  color: white;
  width: 100%;
}
.submit-button:hover {
  background: #15803d;
}

.form-group {
  margin-bottom: 1.5rem;
  text-align: left;
}
label {
  display: block;
  margin-bottom: 0.4rem;
  font-weight: 600;
}
.form-input {
  width: 90%;
  padding: 0.75rem 1rem;
  border-radius: 0.75rem;
  border: 1px solid #cbd5e1;
  font-size: 1rem;
}

.errors {
  margin-top: 1rem;
  list-style: none;
  color: #dc2626;
  text-align: left;
  font-size: 0.9rem;
  padding-left: 1rem;
}

.message {
  margin-top: 1rem;
  color: green;
  font-weight: 600;
}
</style>
