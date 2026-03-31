<script lang="ts">
    let {
        position = "top",
        disabled = false,
        children
    } : {
        position?: "top" | "bottom" | "left" | "right",
        disabled?: boolean,
        children?: () => any
    } = $props();

    let showTooltip = $state(false);

    let tooltipTimeout = $state();

    const handleMouseEnter = () => {
        if (!disabled) {
            tooltipTimeout = setTimeout(() => {
                showTooltip = true;
            }, 300); // Show tooltip after 300ms
        }
    };

    const handleMouseLeave = () => {
        clearTimeout(tooltipTimeout);
        showTooltip = false;
    };

    const handleClick = (e: Event) => {
        e.preventDefault();
        if (disabled) return;
        showTooltip = !showTooltip;
    };

</script>

<div class="tooltip">
    <button class="tooltip-icon {position}" 
        class:disabled={disabled} 
        onmouseenter={handleMouseEnter} 
        onmouseleave={handleMouseLeave}
        onclick={handleClick}
        >
        &#x1F6C8;
    </button>
    {#if showTooltip && !disabled}
        <div class="tooltip-text">
            {#if !children}
                BROKEN
                {console.error("Tooltip component requires children")}
            {:else}
                {@render children()}
            {/if}
        </div>
    {/if}
</div>

<style lang="scss">
    .tooltip {
        position: relative;
        display: inline-block;
        z-index: 3;

        .tooltip-icon {
            background-color: transparent;
            color: #0056b3;
            padding: 0;
            margin: 0 calc($padding / 2);
            border: none;
            font-size: 20px;
            line-height: 20px;
            width: 24px;
            height: 24px;
            border-radius: 50%;
            cursor: pointer;
            display: flex;
            justify-content: center;
            align-items: center;

            &.disabled {
                color: #ccc;
                cursor: not-allowed;
            }

            &:hover:not(.disabled) {
                color: #007bff;
            }
        }

        .tooltip-text {
            visibility: visible;
            width: max-content;
            max-width: 200px;
            background-color: #333;
            color: #fff;
            text-align: left;
            border-radius: $border-radius;
            padding: 0 $padding;
            position: absolute;
            z-index: 1;
            font-size: 14px;

            &.top {
                bottom: 125%;
                //left: 50%;
                transform: translateX(-50%);
                margin-bottom: 5px;
            }

            &.bottom {
                top: 125%;
                left: 50%;
                transform: translateX(-50%);
                margin-top: 5px;
            }

            &.left {
                top: 50%;
                right: 125%;
                transform: translateY(-50%);
                margin-right: 5px;
            }

            &.right {
                top: 50%;
                left: 125%;
                transform: translateY(-50%);
                margin-left: 5px;
            }
        }
    }
</style>