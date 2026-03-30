<script lang="ts">
    import { onMount } from "svelte";
    import Tooltip from "../basic/tooltip.svelte";

    let {
        options = {},
        label,
        value=$bindable(),
        placeholder = "Select an option",
        disabled = false,
        errorMsg = "",
        warningMsg = "",
        bShowValue = false,
        children
    } : {
        options: { 
            [header: string]: { label: string, value: string }[]
        },
        label: string,
        value?: any,
        placeholder?: string,
        disabled?: boolean,
        errorMsg?: string,
        warningMsg?: string,
        bShowValue?: boolean,
        children?: () => any
    } = $props();

    let isOpen = $state(false);
    let selectorElem : HTMLButtonElement | null = null;
    let optionsElem : HTMLDivElement | null = null;

    let btnLabel = $state(placeholder);

    const updateBtnLabel = () => {
        for (const header in options) {
            const option = options[header].find((opt) => opt.value === value);
            if (option) {
                btnLabel = option.label;
                return;
            }
        }
        btnLabel = placeholder;
    }

    onMount(() => {
        if (value && value != "" ) {
            updateBtnLabel();
        }
    });

    const positionOptions = () => {
        setTimeout(() => {
            const selectorRect = selectorElem?.getBoundingClientRect();
            const optionsRect = optionsElem?.getBoundingClientRect();
            if (selectorRect && optionsRect) {
                const viewportHeight = window.innerHeight;
                const spaceBelow = viewportHeight - selectorRect.top;
                const spaceAbove = selectorRect.bottom;
                optionsElem!.style.left = `${selectorRect.left}px`;
                optionsElem!.style.width = `${selectorRect.width}px`;
                if (spaceBelow < optionsRect.height && spaceAbove > optionsRect.height) {
                    optionsElem!.style.top = "auto";
                    optionsElem!.style.bottom = `${viewportHeight - selectorRect.bottom}px`;
                } else {
                    optionsElem!.style.top = `${selectorRect.top}px`;
                    optionsElem!.style.bottom = "auto";
                }
            }
        }, 5);
    }

    const handleOpenSelector = (e : MouseEvent) => {
        e.preventDefault();
        if (disabled) return;
        isOpen = true;
        positionOptions();
        setTimeout(() => {
            const handleClickOutside = (event: MouseEvent) => {
                if (selectorElem && !selectorElem.contains(event.target as Node) &&
                    optionsElem && !optionsElem.contains(event.target as Node)) {
                    isOpen = false;
                    document.removeEventListener('click', handleClickOutside);
                }
            };
            document.addEventListener('click', handleClickOutside);
        }, 0);
    }

    const handleClickOutside = (event: MouseEvent) => {
        if (selectorElem && !selectorElem.contains(event.target as Node) &&
            optionsElem && !optionsElem.contains(event.target as Node)) {
            isOpen = false;
            document.removeEventListener('click', handleClickOutside);
        }
    };


    const handleClickOption = (val: string) => {
        return (e: Event) => {
            e.preventDefault();
            value = val;
            updateBtnLabel();
            document.removeEventListener('click', handleClickOutside);
            isOpen = false;
        }
    }
</script>

{#snippet snippet_option(value: string, label: string)}
 <button class="option" onclick={handleClickOption(value)}>
    <span>{label}</span>
    {#if bShowValue}
        <span class="option-value">{value}</span>
    {/if}
</button>
{/snippet}

<div class="select form-input">
    <div class="input-row">
        <div class="label-row">
            {#if !!children}
                <Tooltip position="top">
                    {@render children()}
                </Tooltip>
            {/if}
            <p class="label" class:error={!!errorMsg} class:warning={!!warningMsg}>
                {label}
            </p>
        </div>
        <button 
            class="selector" 
            onclick={handleOpenSelector} 
            bind:this={selectorElem} 
            disabled={disabled}
            class:error={!!errorMsg}
            class:warning={!!warningMsg}
            >
            <span>{btnLabel}</span>
            {#if bShowValue}
                 <span class="value">{value == "" ? "" : `(${value})`}</span>
            {/if}
            <span class="dropdown-icon">&#9662;</span>
        </button>
    </div>
    {#if !disabled}
        <div class="messages-row">
            <div class="spacer"></div>
            {#if errorMsg}
                <p class="error"><span class="icon">&#x20E0;</span>{errorMsg}</p>
            {:else if warningMsg}
                <p class="warning"><span class="icon">&#x26A0;</span>{warningMsg}</p>
            {/if}
        </div>
    {/if}
</div>
{#if isOpen}
    <div class="select__options" bind:this={optionsElem}>
        <p class="option-label" class:error={!!errorMsg} class:warning={!!warningMsg}>
            {placeholder}
        </p>
        {#if Object.keys(options).length === 0}
            <p class="no-options">No options available</p>
        {:else if Object.keys(options).length === 1}
            {#each options[Object.keys(options)[0]] as option}
                {@render snippet_option(option.value, option.label)}
            {/each}
        {:else}
            {#each Object.keys(options) as optionHeader}
                <h3 class="option-header">{optionHeader}</h3>
                {#each options[optionHeader] as option}
                    {@render snippet_option(option.value, option.label)}
                {/each}
            {/each}
        {/if}
    </div>
{/if}

<style lang="scss">
    .select {
        display: flex;
        align-items: center;
        flex-direction: column;

        .selector {
            flex: 3;
            cursor: pointer;
            margin-right: $padding;
            padding: $padding;            
            border: none;
            border-radius: $border-radius;
            outline: 1px solid #ccc;
            text-align: left;

            display: flex;
            justify-content: space-between;
            align-self: center;
            justify-self: flex-start;
        }

        &__options {
            width: max-content;
            display: flex;
            flex-direction: column;
            border-radius: $border-radius;
            background-color: $clr_secondary;
            max-height: 300px;
            overflow-y: auto;
            position: absolute;
            
            scrollbar-width: thin;
            scrollbar-color: $clr_light transparent;
            &::-webkit-scrollbar {
                width: 8px;
            }
            &::-webkit-scrollbar-track {
                background: transparent;
            }
            &::-webkit-scrollbar-thumb {
                background-color: $clr_light;
                border-radius: $border-radius;
                border: none;
            }

            .option-label {
                background-color: $clr_light;
                color: $clr_dark;
                font-size: 14px;
                border-bottom: 1px solid $clr_light;
                padding: 0;
                padding: $padding;
                margin: 0;
            }
        
            .option-header {
                font-weight: bold;
                border-bottom: 1px solid $clr_light;
                font-size: 16px;
                padding: 0;
                margin: 0;
                color: $clr_light;
            }
        
            .option {
                padding: $padding;
                text-align: left;
                cursor: pointer;
                border-radius: 0;
                outline: none;
                border: none;
                display: flex;
                justify-content: space-between;
                &:hover {
                    background-color: $clr_primary;
                }

                &:last-child {
                    border-radius: 0 0 $border-radius $border-radius;
                }
            }
        }
    }
</style>