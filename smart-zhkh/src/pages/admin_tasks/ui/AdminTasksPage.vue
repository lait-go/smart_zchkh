<template>
  <div class="admin-tasks-page">
    <h2 class="title">📋 Заявки</h2>

    <div class="filters">
      <input
        v-model="searchQuery"
        type="text"
        placeholder="🔍 Поиск по заголовку"
        class="filter-input"
      />

      <select v-model="selectedCategory" class="filter-select">
        <option value="">Все категории</option>
        <option value="plumbing">Сантехника</option>
        <option value="electrical">Электрика</option>
        <option value="cleaning">Уборка</option>
        <option value="elevator">Лифт</option>
        <option value="other">Другое</option>
      </select>

      <select v-model="selectedStatus" class="filter-select">
        <option value="">Все статусы</option>
        <option value="created">Создана</option>
        <option value="in_progress">В работе</option>
        <option value="done">Выполнена</option>
      </select>
    </div>

    <div v-if="loading" class="status">Загрузка...</div>
    <div v-if="error" class="status error">{{ error }}</div>

    <table v-if="filteredTasks.length" class="tasks-table">
      <thead>
        <tr>
          <th>ID</th>
          <th>Категория</th>
          <th>Заголовок</th>
          <th>Статус</th>
          <th>Дата</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="task in filteredTasks" :key="task.id">
          <td>{{ task.id }}</td>
          <td>{{ task.category }}</td>
          <td>
            <router-link :to="`/admin/tasks/${task.id}`" class="task-link">
              {{ task.title }}
            </router-link>
          </td>
          <td>
            <span :class="`status-badge ${task.status}`">
              {{
                task.status === 'created'
                  ? 'Создана'
                  : task.status === 'in_progress'
                    ? 'В работе'
                    : 'Выполнена'
              }}
            </span>
          </td>
          <td>{{ formatDate(task.created_at) }}</td>
        </tr>
      </tbody>
    </table>

    <p v-if="!loading && !filteredTasks.length" class="status">Заявки не найдены</p>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';

const tasks = ref<any[]>([]);
const loading = ref(true);
const error = ref('');

const searchQuery = ref('');
const selectedCategory = ref('');
const selectedStatus = ref('');

async function loadTasks() {
  try {
    const res = await fetch('http://localhost:8081/api/v1/tasks');
    if (!res.ok) throw new Error('Ошибка загрузки');
    tasks.value = await res.json();
  } catch (e) {
    error.value = 'Не удалось загрузить заявки';
  } finally {
    loading.value = false;
  }
}

function formatDate(dateStr: string) {
  return new Date(dateStr).toLocaleString();
}

const filteredTasks = computed(() => {
  return tasks.value.filter((task) => {
    const matchesTitle = task.title.toLowerCase().includes(searchQuery.value.toLowerCase());
    const matchesCategory =
      selectedCategory.value === '' || task.category === selectedCategory.value;
    const matchesStatus = selectedStatus.value === '' || task.status === selectedStatus.value;
    return matchesTitle && matchesCategory && matchesStatus;
  });
});

onMounted(loadTasks);
</script>

<style scoped>
.admin-tasks-page {
  background: white;
  padding: 2rem;
  border-radius: 1rem;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
}

.title {
  font-size: 1.75rem;
  font-weight: bold;
  margin-bottom: 1.5rem;
}

.filters {
  display: flex;
  gap: 1rem;
  margin-bottom: 1.5rem;
  flex-wrap: wrap;
}

.filter-input,
.filter-select {
  padding: 0.5rem 1rem;
  border: 1px solid #ccc;
  border-radius: 0.5rem;
  font-size: 1rem;
}

.tasks-table {
  width: 100%;
  border-collapse: collapse;
}

.tasks-table th,
.tasks-table td {
  padding: 0.75rem;
  border: 1px solid #ddd;
  text-align: left;
}

.status {
  margin-top: 1rem;
  font-weight: 500;
}

.status.error {
  color: red;
}

.task-link {
  color: #2563eb;
  text-decoration: underline;
}
.task-link:hover {
  text-decoration: none;
}

.status-badge {
  padding: 0.25rem 0.5rem;
  border-radius: 0.5rem;
  font-weight: 600;
  text-transform: capitalize;
}
.status-badge.created {
  background: #fef3c7;
  color: #b45309;
}
.status-badge.in_progress {
  background: #e0f2fe;
  color: #0369a1;
}
.status-badge.done {
  background: #dcfce7;
  color: #15803d;
}
</style>
