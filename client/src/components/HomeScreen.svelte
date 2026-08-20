<script lang="ts">
	import MenuSection from './MenuSection.svelte';
	import MenuItemCard from './MenuItemCard.svelte';
	import CartFooterBar from './CartFooterBar.svelte';

	type MenuItem = {
		id: string;
		name: string;
		price: number;
		image?: string | null;
		quantity: number;
	};

	// Swap with real data — this is just enough to preview the layout.
	let sections = $state<{ title: string; items: MenuItem[] }[]>([
		{
			title: 'เมนูแนะนำ',
			items: [{ id: 'r1', name: 'เกี๊ยวเตี๋ยวเรือหมู', price: 15, image: null, quantity: 0 }]
		},
		{
			title: 'ก๋วยเตี๋ยว',
			items: [
				{ id: 'n1', name: 'เส้นเล็กน้ำตก', price: 15, image: null, quantity: 1 },
				{ id: 'n2', name: 'เส้นใหญ่แห้ง', price: 15, image: null, quantity: 0 }
			]
		},
		{
			title: 'เครื่องดื่ม',
			items: [{ id: 'd1', name: 'ชาไทยเย็น', price: 20, image: null, quantity: 0 }]
		}
	]);

	let itemCount = $derived(
		sections.flatMap((s) => s.items).reduce((sum, i) => sum + i.quantity, 0)
	);
	let total = $derived(
		sections.flatMap((s) => s.items).reduce((sum, i) => sum + i.quantity * i.price, 0)
	);

	let { onopencart }: { onopencart?: () => void } = $props();
</script>

<div class="home-screen">
	<header class="home-screen__header">
		<h1>ก๋วยเตี๋ยวเรือ</h1>
	</header>

	<div class="home-screen__scroll">
		{#each sections as section (section.title)}
			<MenuSection title={section.title}>
				{#each section.items as item (item.id)}
					<MenuItemCard
						name={item.name}
						price={item.price}
						image={item.image}
						bind:quantity={item.quantity}
					/>
				{/each}
			</MenuSection>
		{/each}
	</div>

	<footer class="home-screen__footer">
		<CartFooterBar {itemCount} {total} onclick={onopencart} />
	</footer>
</div>
