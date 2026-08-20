<script lang="ts">
	import QuantityStepper from './QuantityStepper.svelte';

	let {
		name,
		image = null,
		quantity = $bindable(1)
	}: {
		name: string;
		image?: string | null;
		quantity?: number;
	} = $props();
</script>

<div class="cart-preview">
	<div class="cart-preview__image">
		{#if image}
			<img src={image} alt={name} loading="lazy" />
		{/if}
	</div>
	<div class="cart-preview__info">
		<p class="cart-preview__name">{name}</p>
		<QuantityStepper bind:value={quantity} />
	</div>
</div>
