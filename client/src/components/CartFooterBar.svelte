<script lang="ts">
	import { Button } from 'bits-ui';
	import { ShoppingBasket } from 'lucide-svelte';

	let {
		itemCount,
		total,
		onclick
	}: {
		itemCount: number;
		total: number;
		onclick?: () => void;
	} = $props();
</script>

<Button.Root class="cart-bar" {onclick} disabled={itemCount === 0}>
	<span class="cart-bar__count">{itemCount}</span>
	<span class="cart-bar__label">
		<ShoppingBasket size={18} />
		ตะกร้า
	</span>
	<span class="cart-bar__total">{total.toLocaleString('th-TH')} บาท</span>
</Button.Root>
