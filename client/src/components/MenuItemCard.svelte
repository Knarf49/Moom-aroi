<script lang="ts">
	import { Button } from 'bits-ui';
	import { Plus } from 'lucide-svelte';
	import QuantityStepper from './QuantityStepper.svelte';

	let {
		name,
		price,
		image = null,
		quantity = $bindable(0),
		onadd
	}: {
		name: string;
		price: number;
		image?: string | null;
		quantity?: number;
		onadd?: () => void;
	} = $props();

	function addFirst() {
		quantity = 1;
		onadd?.();
	}
</script>

<div class="menu-item">
	<div class="menu-item__image">
		{#if image}
			<img src={image} alt={name} loading="lazy" />
		{/if}
	</div>

	<div class="menu-item__info">
		<p class="menu-item__name">{name}</p>
		<p class="menu-item__price">{price} บาท</p>
	</div>

	<div class="menu-item__action">
		{#if quantity > 0}
			<QuantityStepper bind:value={quantity} size="sm" />
		{:else}
			<Button.Root class="menu-item__add" onclick={addFirst} aria-label="เพิ่ม {name} ลงตะกร้า">
				<Plus size={18} />
			</Button.Root>
		{/if}
	</div>
</div>
