<script lang="ts">
  import { PUBLIC_API_BASE } from '$env/static/public';
  import MenuItemCard from '../components/MenuItemCard.svelte';
  import type { PageData } from './$types';

  let { data }: { data: PageData } = $props();
</script>

<main>	
	<div class="home-screen">
		<header class="home-screen__header">
			<h1>ก๋วยเตี๋ยวเรือ</h1>
		</header>

		<div class="home-screen__scroll">
      <h1>เมนูแนะนำ</h1>
			{#each data.menus as item (item.ID)} 
				<MenuItemCard
					name={item.Title}
					price={item.Price}
					image={item.Url ? PUBLIC_API_BASE + item.Url : null}
				/>
			{/each}
      <h1>ก๋วยเตี๋ยว</h1>
      <h1>เครื่องดื่ม</h1>
		</div>
	</div>
</main>
