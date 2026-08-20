<script lang="ts">
	import StatusBadge from './StatusBadge.svelte';

	let {
		name,
		image = null,
		status
	}: {
		name: string;
		image?: string | null;
		status: 'pending' | 'served' | 'cancelled';
	} = $props();
</script>

<div class="order-row">
	<div class="order-row__image">
		{#if image}
			<img src={image} alt={name} loading="lazy" />
		{/if}
	</div>
	<p class="order-row__name">{name}</p>
	<StatusBadge {status} />
</div>
