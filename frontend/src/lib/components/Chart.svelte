<script lang="ts">
    import { onMount, onDestroy } from "svelte"
    import { Chart as ChartJS } from "chart.js"
    import {
        Colors,
        BarController,
        CategoryScale,
        LineController,
        LinearScale,
        PointElement,
        LineElement,
        BarElement,
        Legend,
    } from "chart.js"

    let { type, data, options = {}, plugins = [], ...props } = $props()

    //Register graph types that can be used
    ChartJS.register(
        Colors,
        BarController,
        BarElement,
        CategoryScale,
        LineController,
        LinearScale,
        PointElement,
        LineElement,
        Legend,
    )

    let canvasRef: HTMLCanvasElement
    let chart: ChartJS | null = null

    onMount(() => {
        chart = new ChartJS(canvasRef, {
            type,
            data,
            options,
            plugins,
        })
    })

    $effect(() => {
        if (!chart) return

        chart.data = data
        Object.assign(chart.options, options)
        chart.update()
    })

    onDestroy(() => {
        if (chart) chart.destroy()
        chart = null
    })
</script>

<canvas
    bind:this={canvasRef}
    {...props}
>
</canvas>
