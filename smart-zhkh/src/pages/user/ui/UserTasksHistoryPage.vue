<template>
  <div class="user-history-page">
    <h2>📜 История заявок</h2>

    <div v-if="loading" class="status">Загрузка...</div>
    <div v-if="error" class="status error">{{ error }}</div>

    <div v-for="task in tasks" :key="task.id" class="task-block">
      <div class="task-info">
        <p><strong>ID:</strong> {{ task.id }}</p>
        <p><strong>Категория:</strong> {{ task.category }}</p>
        <p><strong>Заголовок:</strong> {{ task.title }}</p>
        <p><strong>Статус:</strong> {{ formatStatus(task.status) }}</p>
        <p><strong>Создана:</strong> {{ formatDate(task.created_at) }}</p>
      </div>

      <div class="comment-section">
        <h4>💬 Комментарии:</h4>
        <ul v-if="task.comments && task.comments.length" class="comment-list">
          <li v-for="(comment, index) in task.comments" :key="index">
            <strong>{{ comment.user }}:</strong> {{ comment.text }}
          </li>
        </ul>
        <p v-else class="no-comments">Комментариев нет</p>

        <textarea
          v-model="task.newComment"
          class="comment-input"
          placeholder="Оставьте комментарий для администратора..."
        ></textarea>
        <button class="submit-comment" @click="submitComment(task)">📨 Отправить</button>
      </div>
    </div>

    <p v-if="!loading && !tasks.length" class="status">Заявок не найдено</p>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useAuthStore } from '@/shared/store/auth';

const auth = useAuthStore();
const tasks = ref<any[]>([]);
const loading = ref(true);
const error = ref('');

async function loadMyTasks() {
  try {
    const res = await fetch(`http://localhost:8081/api/v1/tasks?account_id=${auth.userId}`);
    if (!res.ok) throw new Error();
    const data = await res.json();

    // добавим поля для комментариев
    for (const task of data) {
      task.comments = await loadComments(task.id);
      task.newComment = '';
    }

    tasks.value = data;
  } catch {
    error.value = '❌ Не удалось загрузить заявки';
  } finally {
    loading.value = false;
  }
}

function formatDate(str: string) {
  return new Date(str).toLocaleString();
}

function formatStatus(status: string) {
  switch (status) {
    case 'created':
      return 'Создана';
    case 'in_progress':
      return 'В работе';
    case 'done':
      return 'Завершена';
    default:
      return status;
  }
}

async function loadComments(taskId: string) {
  try {
    const res = await fetch(`http://localhost:8081/api/v1/tasks/comment/${taskId}`);
    if (!res.ok) return [];
    return await res.json();
  } catch {
    return [];
  }
}

async function submitComment(task: any) {
  if (!task.newComment.trim()) return;

  const comment = {
    task_id: task.id,
    user: auth.username,
    text: task.newComment,
  };

  try {
    const res = await fetch('http://localhost:8081/api/v1/tasks/comment', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(comment),
    });
    if (!res.ok) throw new Error();

    task.comments.push({ user: auth.username, text: task.newComment });
    task.newComment = '';
  } catch {
    alert('❌ Не удалось отправить комментарий');
  }
}

onMounted(loadMyTasks);
</script>

<style scoped>
.user-history-page {
  background: white;
  padding: 2rem;
  border-radius: 1rem;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
  font-family: 'Segoe UI', sans-serif;
}

h2 {
  font-size: 1.5rem;
  font-weight: bold;
  margin-bottom: 1.5rem;
}

.task-block {
  border: 1px solid #ddd;
  border-radius: 1rem;
  padding: 1rem 1.5rem;
  margin-bottom: 2rem;
  background: #f9fafb;
}

.task-info p {
  margin: 0.2rem 0;
}

.comment-section {
  margin-top: 1rem;
}

.comment-list {
  margin-bottom: 0.5rem;
  padding-left: 1.2rem;
}

.comment-list li {
  margin-bottom: 0.3rem;
}

.no-comments {
  font-style: italic;
  color: #888;
}

.comment-input {
  width: 100%;
  min-height: 60px;
  padding: 0.5rem;
  margin-top: 0.5rem;
  border-radius: 0.5rem;
  border: 1px solid #ccc;
  font-size: 1rem;
}

.submit-comment {
  margin-top: 0.5rem;
  padding: 0.4rem 1rem;
  background: #2563eb;
  color: white;
  border: none;
  border-radius: 0.5rem;
  cursor: pointer;
  font-weight: 600;
}
.submit-comment:hover {
  background: #1e40af;
}

.status {
  margin-top: 1rem;
  font-weight: 500;
}

.status.error {
  color: red;
}
</style>
