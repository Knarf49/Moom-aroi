import { PUBLIC_API_BASE } from '$env/static/public';
import type { Dish } from '$lib/types.js';

export async function load({ fetch }) {
	const res = await fetch(`${PUBLIC_API_BASE}/api/menu`);

	if (!res.ok) {
		throw new Error('Failed to fetch menu');
	}

	const menus: Dish[] = await res.json();
	return { menus };
}
