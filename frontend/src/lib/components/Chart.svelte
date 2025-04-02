<script lang="ts">
    import {
        BarController,
        BarElement,
        CategoryScale,
        Chart as ChartJS,
        Colors,
        Legend,
        LineController,
        LineElement,
        LinearScale,
        PointElement,
        type ChartConfiguration,
        type ChartType,
        type DefaultDataPoint,
    } from "chart.js"
    import { onDestroy, onMount } from "svelte"

    type TType = ChartType
    type TData = DefaultDataPoint<TType>
    type TLabel = unknown

    let {
        type,
        data,
        options = {},
        plugins = [],
        ...props
    }: ChartConfiguration<TType, TData, TLabel> = $props()

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
