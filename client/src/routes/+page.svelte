<script lang="ts">
	import { onMount } from 'svelte';
    import { PUBLIC_API_BASE, PUBLIC_WS_URL } from '$env/static/public';

	let httpStatus = $state('Pending...');
	let wsStatus = $state('Connecting... 🟡');
	let messages = $state<string[]>([]);
	let ws: WebSocket | undefined;

	const API_BASE = PUBLIC_API_BASE;
	const WS_URL = PUBLIC_WS_URL;

	onMount(() => {
		// --- 1. Test HTTP Connection ---
		(async () => {
			try {
				const res = await fetch(`${API_BASE}/api/menu`);
				if (res.ok) {
					const data = await res.json();
					httpStatus = `Connected 🟢 (Response: ${JSON.stringify(data)})`;
				} else {
					httpStatus = `Failed 🔴 (Status: ${res.status})`;
				}
			} catch (error) {
				httpStatus = `Error 🔴 (${error instanceof Error ? error.message : String(error)})`;
			}
		})();

		// --- 2. Test WebSocket Connection ---
		ws = new WebSocket(WS_URL);

		ws.onopen = () => {
			wsStatus = 'Connected 🟢';

			const testPayload = {
				event: 'toggle_menu_status',
				payload: {
					menuId: 'dish_123',
					newStatus: 0
				}
			};

			ws?.send(JSON.stringify(testPayload));
			messages = [...messages, `Sent: ${JSON.stringify(testPayload)}`];
		};

		ws.onmessage = (event: MessageEvent) => {
			messages = [...messages, `Received: ${event.data}`];
		};

		ws.onerror = (error: Event) => {
			wsStatus = 'Error 🔴 (Check console)';
			console.error('WebSocket Error:', error);
		};

		ws.onclose = () => {
			wsStatus = 'Disconnected 🔴';
		};

		// Cleanup the WebSocket when the component unmounts
		return () => {
			ws?.close();
		};
	});
</script>

<main class="p-8 font-sans max-w-2xl mx-auto">
	<h1 class="text-2xl font-bold mb-6">Backend Connection Diagnostic</h1>

	<div class="bg-gray-100 p-4 rounded-lg mb-6 shadow">
		<h2 class="font-bold text-lg mb-2">HTTP API Route (GET /api/menu)</h2>
		<p class="font-mono text-sm">{httpStatus}</p>
	</div>

	<div class="bg-gray-100 p-4 rounded-lg shadow">
		<h2 class="font-bold text-lg mb-2">WebSocket ({WS_URL})</h2>
		<p class="font-mono text-sm mb-4">Status: {wsStatus}</p>

		<h3 class="font-bold text-sm text-gray-600 mb-2">Message Log:</h3>
		<div class="bg-black text-green-400 p-3 rounded font-mono text-xs h-48 overflow-y-auto">
			{#if messages.length === 0}
				<p class="text-gray-500">Waiting for messages...</p>
			{/if}
			{#each messages as msg}
				<p class="mb-1">&gt; {msg}</p>
			{/each}
		</div>
	</div>
</main>