<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useAuthStore } from '@/shared/store/auth';
import { fetchCharges, updateCharge } from '@/shared/api/charges';

const auth = useAuthStore();
const charges = ref<any[]>([]);
const isLoading = ref(true);

const showModal = ref(false);
const selectedChargeId = ref<number | null>(null);
const cardNumber = ref('');
const cvv = ref('');

onMounted(async () => {
  try {
    charges.value = await fetchCharges(auth.userId);
  } catch (e) {
    console.error('Ошибка загрузки начислений:', e);
  } finally {
    isLoading.value = false;
  }
});

const unpaidCharges = computed(() => charges.value.filter((c) => !c.paid));
const paidCharges = computed(() => charges.value.filter((c) => c.paid));

function formatDate(date: string) {
  return new Date(date).toLocaleDateString('ru-RU');
}

function openPaymentModal(id: number) {
  selectedChargeId.value = id;
  cardNumber.value = '';
  cvv.value = '';
  showModal.value = true;
}

async function confirmPayment() {
  const id = selectedChargeId.value;
  if (!id || !cardNumber.value || !cvv.value) return;

  const charge = charges.value.find((c) => c.id === id);
  if (!charge) return;

  try {
    await updateCharge({ ...charge, paid: true });
    charge.paid = true;
    showModal.value = false;
  } catch (e) {
    console.error('Ошибка оплаты:', e);
  }
}
</script>

<template>
  <div class="charges-container container">
    <div class="charges-header">
      <h2 class="charges-title">Ваши начисления</h2>
      <RouterLink to="/charges/add" class="btn btn-primary add-button">+ Добавить</RouterLink>
    </div>

    <div v-if="isLoading" class="loading">⏳ Загрузка...</div>

    <div v-else>
      <div v-if="unpaidCharges.length > 0">
        <h3>Неоплаченные</h3>
        <div class="card-grid">
          <div v-for="charge in unpaidCharges" :key="charge.id" class="charge-card">
            <h3 class="category">{{ charge.category }}</h3>
            <p class="amount">💰 {{ charge.amount }} ₽</p>
            <p class="date">📅 {{ formatDate(charge.date) }}</p>
            <button @click="openPaymentModal(charge.id)" class="pay-button">💳 Оплатить</button>
          </div>
        </div>
      </div>

      <div v-if="paidCharges.length > 0" style="margin-top: 2rem">
        <h3>Оплаченные</h3>
        <div class="card-grid">
          <div v-for="charge in paidCharges" :key="charge.id" class="charge-card paid">
            <h3 class="category">{{ charge.category }}</h3>
            <p class="amount">💰 {{ charge.amount }} ₽</p>
            <p class="date">📅 {{ formatDate(charge.date) }}</p>
            <p class="status">✅ Оплачено</p>
          </div>
        </div>
      </div>

      <div v-if="unpaidCharges.length === 0 && paidCharges.length === 0" class="empty-state">
        <p>Начислений нет</p>
      </div>
    </div>
  </div>

  <!-- 💳 Модальное окно оплаты -->
  <teleport to="body">
    <div v-if="showModal" class="modal-overlay">
      <div class="modal-content">
        <h3>Оплата начисления</h3>
        <label>Номер карты</label>
        <input v-model="cardNumber" type="text" placeholder="0000 0000 0000 0000" maxlength="19" />
        <label>CVV</label>
        <input v-model="cvv" type="text" placeholder="123" maxlength="3" />

        <div class="modal-actions">
          <button class="confirm" @click="confirmPayment">Подтвердить</button>
          <button class="cancel" @click="showModal = false">Отмена</button>
        </div>
      </div>
    </div>
  </teleport>
</template>

<style scoped>
.charges-container {
  position: relative;
  overflow: hidden;
  margin: 3rem auto;
  padding: 3rem 2rem;
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(12px);
  border-radius: 1.75rem;
  box-shadow:
    0 12px 30px -8px rgba(0, 0, 0, 0.2),
    inset 0 2px 4px rgba(255, 255, 255, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.4);
  animation: cardEntrance 0.8s cubic-bezier(0.22, 1, 0.36, 1);
  color: var(--color-text-dark);
}
.charges-container::before {
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

.charges-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}
.charges-title {
  font-size: 2rem;
  font-weight: 800;
  position: relative;
  display: inline-block;
}
.charges-title::after {
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

.add-button {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
}
.btn-icon {
  width: 1.25rem;
  height: 1.25rem;
  stroke: currentColor;
  stroke-width: 2;
  fill: none;
  transition: transform 0.3s;
}
.add-button:hover .btn-icon {
  transform: scale(1.2);
}

.loading,
.empty-state {
  text-align: center;
  padding: 2rem 0;
  font-size: 1rem;
  color: var(--color-text-dark);
}

.card-grid {
  display: grid;
  gap: 2rem;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
}

.charge-card {
  position: relative;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(8px);
  border-radius: 1.25rem;
  box-shadow:
    0 6px 16px -4px rgba(0, 0, 0, 0.1),
    inset 0 1px 3px rgba(255, 255, 255, 0.5);
  padding: 2rem 1.5rem 1.5rem;
  transition:
    transform 0.4s cubic-bezier(0.22, 1, 0.36, 1),
    box-shadow 0.4s;
  overflow: hidden;
  color: var(--color-text-dark);
  text-align: center;
}
.charge-card:hover {
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

.charge-icon {
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
.charge-card:hover .charge-icon {
  transform: scale(1.2) translateY(-4px);
  filter: drop-shadow(0 4px 8px rgba(37, 99, 235, 0.3));
}

.category {
  font-size: 1.25rem;
  font-weight: 700;
  color: var(--color-primary-dark);
  margin-bottom: 0.5rem;
}

.amount {
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--color-success);
  margin-bottom: 0.5rem;
}

.date {
  font-size: 0.875rem;
  color: var(--color-text-dark);
}

.success-message {
  margin-bottom: 1rem;
  color: #16a34a;
  font-weight: 500;
}

.pay-button {
  background: #22c55e;
  border: none;
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 6px;
  cursor: pointer;
  margin-top: 1rem;
}
.pay-button:hover {
  background: #16a34a;
}

.confirm {
  background: #2563eb;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 6px;
}
.cancel {
  background: transparent;
  color: #64748b;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 6px;
}

.charges-container {
  max-width: 800px;
  margin: 2rem auto;
  padding: 2rem;
  background: white;
  border-radius: 12px;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.04);
}

.charges-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 1rem;
}

.charge-card {
  background: #f9fafb;
  border-radius: 8px;
  padding: 1rem;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.05);
  text-align: center;
}
.charge-card.paid {
  background: #e0fce0;
  border: 1px solid #22c55e;
}

.category {
  font-weight: bold;
  font-size: 1.2rem;
}

.amount,
.date {
  margin: 0.25rem 0;
}

.status {
  color: #16a34a;
  font-weight: 500;
}

.pay-button {
  margin-top: 0.5rem;
  padding: 0.5rem 1rem;
  background: #3b82f6;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
}
.pay-button:hover {
  background: #2563eb;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.4);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 100;
}
.modal-content {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  max-width: 400px;
  width: 90%;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.2);
}
.modal-content h3 {
  margin-bottom: 1rem;
}
.modal-content input {
  width: 100%;
  margin-bottom: 1rem;
  padding: 0.5rem;
  border-radius: 6px;
  border: 1px solid #cbd5e1;
}
.modal-actions {
  display: flex;
  justify-content: space-between;
}
.confirm {
  background: #2563eb;
  color: white;
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 6px;
  cursor: pointer;
}
.cancel {
  background: transparent;
  color: #64748b;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 6px;
  cursor: pointer;
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

/* Responsive */
@media (max-width: 768px) {
  .charges-container {
    padding: 2rem 1rem;
  }
  .charges-title {
    font-size: 1.5rem;
  }
  .card-grid {
    gap: 1.5rem;
  }
}
@media (max-width: 480px) {
  .charge-card {
    padding: 1.5rem 1rem 1rem;
  }
  .category {
    font-size: 1.125rem;
  }
  .amount {
    font-size: 1rem;
  }
  .date {
    font-size: 0.8rem;
  }
}
</style>