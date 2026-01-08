export async function createTask(formData: FormData) {
  const response = await fetch('http://localhost:8081/api/v1/tasks', {
    method: 'POST',
    body: formData,
  });

  if (!response.ok) {
    const error = await response.json().catch(() => ({}));
    throw new Error(error.message || 'Ошибка при создании заявки');
  }

  return await response.json();
}
