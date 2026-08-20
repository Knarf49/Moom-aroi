<script lang="ts">
	import { Button } from 'bits-ui';
	import { Minus, Plus } from 'lucide-svelte';

	let {
		value = $bindable(1),
		min = 0,
		max = 99,
		size = 'md', // "sm" | "md"
		onchange
	}: {
		value?: number;
		min?: number;
		max?: number;
		size?: 'sm' | 'md';
		onchange?: (value: number) => void;
	} = $props();

	function dec() {
		if (value <= min) return;
		value -= 1;
		onchange?.(value);
	}

	function inc() {
		if (value >= max) return;
		value += 1;
		onchange?.(value);
	}
</script>

<div class="stepper stepper--{size}">
	<Button.Root class="stepper__btn" onclick={dec} disabled={value <= min} aria-label="ลดจำนวน">
		<Minus size={size === 'sm' ? 14 : 16} />
	</Button.Root>

	<span class="stepper__value" aria-live="polite">{value}</span>

	<Button.Root class="stepper__btn" onclick={inc} disabled={value >= max} aria-label="เพิ่มจำนวน">
		<Plus size={size === 'sm' ? 14 : 16} />
	</Button.Root>
</div>
