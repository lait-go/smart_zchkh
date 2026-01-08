<template>
  <div class="admin-charges container">
    <h2>📋 Управление начислениями</h2>

    <div class="filters">
      <input v-model="filters.account" placeholder="Поиск по счёту" />
      <select v-model="filters.category">
        <option value="">Все категории</option>
        <option value="Вода">Вода</option>
        <option value="Электричество">Электричество</option>
        <option value="Газ">Газ</option>
      </select>
      <input type="date" v-model="filters.date" />
    </div>

    <button @click="openMassChargeModal">➕ Массовое начисление</button>

    <table class="charges-table">
      <thead>
        <tr>
          <th>ФИО</th>
          <th>Счёт</th>
          <th>Категория</th>
          <th>Сумма</th>
          <th>Дата</th>
          <th>Действия</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="charge in filteredCharges" :key="charge.id">
          <td>{{ charge.account_id }}</td>
          <td>{{ charge.category }}</td>
          <td>{{ charge.amount }}</td>
          <td>{{ formatDate(charge.date) }}</td>
          <td>
            <button @click="editCharge(charge)">✏️</button>
            <button @click="viewHistory(charge)">🕘</button>
          </td>
        </tr>
      </tbody>
    </table>

    <!-- Временно для отладки -->
    <pre v-if="debug" style="margin-top: 2rem">{{ filteredCharges }}</pre>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';

const charges = ref<any[]>([]);
const filters = ref({ account: '', category: '', date: '' });
const debug = false;

async function fetchCharges() {
  try {
    const res = await fetch('http://localhost:7070/api/v1/charges');
    charges.value = await res.json();
  } catch (e) {
    console.error('Ошибка загрузки начислений:', e);
  }
}

const filteredCharges = computed(() => {
  return charges.value.filter((c) => {
    const matchAccount = filters.value.account
      ? c.account_id.includes(filters.value.account)
      : true;
    const matchCategory = filters.value.category ? c.category === filters.value.category : true;
    const matchDate = filters.value.date ? c.date.startsWith(filters.value.date) : true;
    return matchAccount && matchCategory && matchDate;
  });
});

function formatDate(date: string) {
  return new Date(date).toLocaleDateString('ru-RU');
}

function editCharge(charge: any) {
  alert(`Редактирование: ${JSON.stringify(charge)}`);
}

function viewHistory(charge: any) {
  alert(`История: ${JSON.stringify(charge.history || [])}`);
}

function openMassChargeModal() {
  alert('Открытие массового добавления...');
}

onMounted(() => {
  fetchCharges();
});
</script>

<style scoped>
.admin-charges {
  max-width: 900px;
  margin: 2rem auto;
  padding: 2rem;
  background: white;
  border-radius: 1rem;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
}

.filters {
  display: flex;
  gap: 1rem;
  margin-bottom: 1.5rem;
}

.filters input,
.filters select {
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 0.5rem;
}

.charges-table {
  width: 100%;
  border-collapse: collapse;
}

.charges-table th,
.charges-table td {
  border: 1px solid #ddd;
  padding: 0.75rem;
  text-align: left;
}

.charges-table th {
  background-color: #f3f4f6;
}

.charges-table button {
  margin-right: 0.5rem;
}
</style>
