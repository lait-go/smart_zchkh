export async function createCharge(charge: {
  user_id: string;
  category: string;
  amount: number;
  date: string;
}) {
  console.log(charge);
  const res = await fetch('http://localhost:7070/api/charges', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(charge),
  });

  if (!res.ok) {
    throw new Error('Ошибка создания начисления');
  }

  return await res.json();
}

export async function fetchCharges(userId: string) {
  const res = await fetch(`http://localhost:7070/api/charges?user_id=${userId}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!res.ok) {
    throw new Error('Не удалось загрузить начисления');
  }

  return await res.json();
}

export async function deleteCharge(id: number) {
  const res = await fetch(`http://localhost:7070/api/charges/${id}`, {
    method: 'DELETE',
  });

  if (!res.ok) {
    throw new Error('Не удалось удалить начисление');
  }

  return await res.json();
}

export async function markChargeAsPaid(id: number) {
  const res = await fetch(`http://localhost:7070/api/charges/${id}/pay`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
  });

  if (!res.ok) {
    throw new Error('Не удалось отметить как оплачено');
  }
}
export async function updateCharge(charge: any) {
  const res = await fetch(`http://localhost:7070/api/charges/${charge.id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(charge),
  });

  if (!res.ok) {
    throw new Error('Ошибка обновления начисления');
  }

  return await res.json();
}
