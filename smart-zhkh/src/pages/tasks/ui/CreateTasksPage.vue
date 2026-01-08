<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/shared/store/auth';
import { createTask } from '@/shared/api/tasks';

const router = useRouter();
const auth = useAuthStore();

const title = ref('');
const description = ref('');
const category = ref('');
const files = ref<File[]>([]);
const message = ref('');

const accountId = computed(() => auth.accountId || 'user-001'); // заглушка если нет ID

function handleFiles(event: Event) {
  const input = event.target as HTMLInputElement;
  if (input.files) {
    files.value = Array.from(input.files);
  }
}

async function submitForm() {
  try {
    const formData = new FormData();
    formData.append('title', title.value);
    formData.append('description', description.value);
    formData.append('category', category.value);
    formData.append('account_id', accountId.value);
    files.value.forEach((file) => formData.append('attachments', file));

    await createTask(formData);

    message.value = '✅ Заявка успешно отправлена!';
    setTimeout(() => router.push('/dashboard'), 1000);
  } catch (err: any) {
    message.value = '❌ ' + (err.message || 'Ошибка при отправке');
  }
}
</script>

<template>
  <div class="add-charge-container container">
    <h2 class="add-charge-title">
      <svg viewBox="0 0 24 24" class="add-icon" aria-hidden="true">
        <line x1="12" y1="5" x2="12" y2="19" />
        <line x1="5" y1="12" x2="19" y2="12" />
      </svg>
      Создать новую заявку
    </h2>

    <form @submit.prevent="submitForm" class="add-charge-form">
      <div class="form-group">
        <label for="title">📝 Заголовок</label>
        <input
          v-model="title"
          id="title"
          type="text"
          required
          placeholder="Например: Течёт кран на кухне"
          class="form-input"
        />
      </div>

      <div class="form-group">
        <label for="description">📋 Описание</label>
        <textarea
          v-model="description"
          id="description"
          rows="4"
          required
          placeholder="Подробности проблемы..."
          class="form-input"
        ></textarea>
      </div>

      <div class="form-group">
        <label for="category">📂 Категория</label>
        <select v-model="category" id="category" required class="form-input">
          <option disabled value="">Выберите категорию</option>
          <option value="plumbing">Сантехника</option>
          <option value="electrical">Электрика</option>
          <option value="cleaning">Уборка</option>
          <option value="elevator">Лифт</option>
          <option value="other">Другое</option>
        </select>
      </div>

      <div class="form-group">
        <label for="files">📎 Прикрепить файлы</label>
        <input id="files" type="file" multiple @change="handleFiles" class="form-input" />
      </div>

      <button type="submit" class="submit-button">Отправить</button>
    </form>

    <p v-if="message" class="message">{{ message }}</p>
  </div>
</template>

<style scoped>
/* Стили аналогичны AddChargesPage.vue */
.add-charge-container {
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
.add-charge-container::before {
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
.add-charge-title {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-size: 1.75rem;
  font-weight: 800;
  color: var(--color-primary-dark);
  margin-bottom: 2rem;
  position: relative;
}
.add-charge-title::after {
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
.add-icon {
  width: 1.75rem;
  height: 1.75rem;
  stroke: var(--color-primary);
  stroke-width: 2;
  fill: none;
  transition: transform 0.3s;
}
.add-icon:hover {
  transform: scale(1.2);
}
.add-charge-form {
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
.message {
  margin-top: 1.5rem;
  font-weight: 600;
  color: var(--color-success);
  text-align: center;
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
  .add-charge-container {
    padding: 2.5rem 1.5rem;
  }
  .add-charge-title {
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
  .add-charge-title {
    font-size: 1.25rem;
  }
  .submit-button {
    font-size: 0.95rem;
  }
}
</style>
