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
    generics="TChartType extends ChartType"
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
    }: DefaultChartConfig<TChartType> = $props()

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

    type TData = DefaultDataPoint<TChartType>
    type TypedChartJS = ChartJS<TChartType, TData, TLabel>
    type ChartOptions = TypedChartJS["options"]

    let canvasRef: HTMLCanvasElement
    let chart: TypedChartJS | null = null

    onMount(() => {
        chart = new ChartJS(canvasRef, {
            type,
            data,
            options: options,
            plugins,
        })
    })

    $effect(() => {
        if (!chart) return

        chart.data = data
        Object.assign(chart.options ?? {}, options)
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
