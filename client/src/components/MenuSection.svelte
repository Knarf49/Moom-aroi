<script lang="ts">
	import type { Snippet } from 'svelte';

	let {
		title,
		children
	}: {
		title: string;
		children: Snippet;
	} = $props();
</script>

<section class="menu-section">
	<h2 class="menu-section__title">{title}</h2>
	<div class="menu-section__list">
		{@render children()}
	</div>
</section>
