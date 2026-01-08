export async function fetchAccounts() {
  const res = await fetch('http://localhost:8080/api/accounts', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!res.ok) throw new Error('Ошибка загрузки счетов');

  return await res.json();
}
export async function fetchAccountById(id: number) {
  const res = await fetch(`http://localhost:8080/api/accounts/${id}`);
  if (!res.ok) throw new Error('Ошибка загрузки счёта');
  return await res.json();
}

export async function updateAccount(account: any) {
  const res = await fetch(`http://localhost:8080/api/accounts/${account.id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(account),
  });

  if (!res.ok) throw new Error('Ошибка обновления счёта');
  return await res.json();
}
export async function createAccount(account: {
  user_id: string;
  account_number: string;
  full_name: string;
  address: string;
  area: number;
}) {
  const res = await fetch('http://localhost:8080/api/accounts', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(account),
  });

  if (!res.ok) {
    const errorText = await res.text();
    throw new Error(`Ошибка создания аккаунта: ${errorText}`);
  }

  return await res.text(); // или res.json(), если backend возвращает JSON
}
