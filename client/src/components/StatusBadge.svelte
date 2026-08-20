<script lang="ts">
	type Status = 'pending' | 'served' | 'cancelled';

	let { status }: { status: Status } = $props();

	const labels: Record<Status, string> = {
		pending: 'กำลังทำ',
		served: 'เสิร์ฟแล้ว',
		cancelled: 'ยกเลิก'
	};
</script>

<span class="status-badge status-badge--{status}">
	{labels[status]}
</span>
