<script lang="ts">
    let { type = 'button', subtype = "", disabled = false, onClick = () => {}, children } : {
        type?: "button" | "submit" | "reset",
        subtype?: string,
        disabled?: boolean,
        onClick?: () => void,
        children?: () => any
    } = $props();

    const handleClick = (e: Event) => {
        e.preventDefault();
        if (disabled) return;
        onClick?.();
    }
</script>

<button class="btn {subtype}" type={type} disabled={disabled} onclick={handleClick}>
    {#if !children}
        BROKEN
        {console.error("Button component requires children")}
    {:else}
        {@render children()}
    {/if}
</button>

<style lang="scss">
    .btn {
        padding: calc($padding * 1.5) calc($padding * 2);
        border: none;
        border-radius: $border-radius;
        cursor: pointer;
        font-size: 18px;
        transition: background-color 0.2s ease;
        background-color: #f0f0f0;
        color: #333;
        font-weight: bold;
        &:hover:not(:disabled) {
            background-color: #e0e0e0;
        }

        &.icon {
            width: 48px;
            height: 48px;
            font-weight: bold;
            display: flex;
            align-items: center;
            justify-content: center;
            background-color: transparent;
            color: #999;
            padding: 0;
            font-size: 32px;

            &.small {
                font-size: 24px;
                width: 32px;
                height: 32px;
            }

            &:hover:not(:disabled) {
                color: #ccc;
                background-color: transparent;
            }
        }

        &.primary {
            background-color: #007bff;
            color: white;

            &:hover:not(:disabled) {
                background-color: #0056b3;
            }
        }

        &.secondary {
            background-color: #6c757d;
            color: white;

            &:hover:not(:disabled) {
                background-color: #5a6268;
            }
        }

        &.add {
            background-color: $clr_success;
            color: white;

            &:hover:not(:disabled) {
                background-color: #218838;
            }
        }

        &.error {
            background-color: $clr_error;
            color: white;

            &:hover:not(:disabled) {
                background-color: #c82333;
            }
        }

        &.list {
            border-radius: 0;
        }

        &.bottom {
            width: 100%;
            border-radius: 0 0 $border-radius $border-radius;
        }

        &:disabled {
            background-color: #e0e0e0;
            color: #a0a0a0;
            cursor: not-allowed;
        }
    }
</style>