<script module>
    type TLabel = string
    export type DefaultChartConfig<TType extends ChartType> = ChartConfiguration<
        TType,
        DefaultDataPoint<TType>,
        TLabel
    >
    export type ChartData<TType extends ChartType> = DefaultChartConfig<TType>["data"]
</script>

<script
    lang="ts"
    generics="TChartType extends ChartType, TChartData = DefaultDataPoint<TChartType>, TChartLabel = unknown"
>
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
        Title,
        type ChartConfiguration,
        type ChartType,
        type DefaultDataPoint,
    } from "chart.js"
    import { onDestroy, onMount } from "svelte"

    let {
        type,
        data,
        options = undefined,
        plugins = [],
        ...props
    }: ChartConfiguration<TChartType, TChartData, TChartLabel> = $props()

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
        Title,
    )

    type TypedChartJS = ChartJS<TChartType, TChartData, TChartLabel>

    let canvasRef: HTMLCanvasElement
    let chart: TypedChartJS | null = null

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
        if (chart.options) Object.assign(chart.options, options)
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
