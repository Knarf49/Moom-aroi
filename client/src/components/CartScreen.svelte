<script lang="ts">
	import { Separator, Button } from 'bits-ui';
	import CartPreviewCard from './CartPreviewCard.svelte';
	import OrderRow from './OrderRow.svelte';

	type OrderedItem = {
		id: string;
		name: string;
		image?: string | null;
		status: 'pending' | 'served' | 'cancelled';
	};

	let {
		previewItemName,
		previewItemImage = null,
		previewQuantity = $bindable(1),
		orderedItems,
		total,
		oncheckout
	}: {
		previewItemName: string;
		previewItemImage?: string | null;
		previewQuantity?: number;
		orderedItems: OrderedItem[];
		total: number;
		oncheckout?: () => void;
	} = $props();
</script>

<div class="cart-screen">
	<header class="cart-screen__header">
		<h1>ตะกร้า</h1>
	</header>

	<div class="cart-screen__scroll">
		<CartPreviewCard
			name={previewItemName}
			image={previewItemImage}
			bind:quantity={previewQuantity}
		/>

		<Separator.Root class="cart-screen__divider" />

		<section>
			<h2 class="cart-screen__section-title">ออร์เดอร์ที่สั่งไปแล้ว</h2>
			<div class="cart-screen__order-list">
				{#each orderedItems as item (item.id)}
					<OrderRow name={item.name} image={item.image} status={item.status} />
				{/each}
			</div>
		</section>
	</div>

	<footer class="cart-screen__footer">
		<Separator.Root class="cart-screen__divider" />
		<div class="cart-screen__total">
			<span>ยอดรวมทั้งหมด</span>
			<span>{total.toLocaleString('th-TH')} บาท</span>
		</div>
		<Button.Root class="cart-screen__checkout" onclick={oncheckout}>คิดเงิน</Button.Root>
	</footer>
</div>
