import {writable} from 'svelte/store'
import {browser} from '$app/environment'

const initialValue = browser ? localStorage.getItem('cart') : ''

export const cart = writable(initialValue)

if(browser){
    cart.subscribe(value =>{
        localStorage.setItem('cart',value as string)
    })
}
