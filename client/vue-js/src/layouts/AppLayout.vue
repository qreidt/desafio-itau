<template>
  <header class="absolute inset-x-0 top-0 z-50 flex h-16 border-b border-gray-900/10">
    <div class="mx-auto flex w-full max-w-7xl items-center justify-between px-4 sm:px-6 lg:px-8">
      <div class="flex flex-1 items-center gap-x-6">
        <button type="button" class="-m-3 p-3 md:hidden" @click="mobileMenuOpen = true">
          <span class="sr-only">Open main menu</span>
          <Bars3Icon class="h-5 w-5 text-gray-900" aria-hidden="true" />
        </button>
        <img class="h-8 w-auto" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600" alt="Your Company" />
      </div>
      <nav class="hidden md:flex md:gap-x-11 md:text-sm md:font-semibold md:leading-6 md:text-gray-700">
        <RouterLink to="/home">Home</RouterLink>
        <RouterLink to="/transactions">Transações</RouterLink>
      </nav>
      <div class="flex flex-1 items-center justify-end gap-x-8">
        <span
          @click="logout"
          class="py-1.5 px-2 rounded bg-gray-100 hover:bg-gray-200 cursor-pointer text-sm"
        >
          Logout
        </span>
      </div>
    </div>
    <Dialog as="div" class="lg:hidden" @close="mobileMenuOpen = false" :open="mobileMenuOpen">
      <div class="fixed inset-0 z-50" />
      <DialogPanel class="fixed inset-y-0 left-0 z-50 w-full overflow-y-auto bg-white px-4 pb-6 sm:max-w-sm sm:px-6 sm:ring-1 sm:ring-gray-900/10">
        <div class="-ml-0.5 flex h-16 items-center gap-x-6">
          <button type="button" class="-m-2.5 p-2.5 text-gray-700" @click="mobileMenuOpen = false">
            <span class="sr-only">Close menu</span>
            <XMarkIcon class="h-6 w-6" aria-hidden="true" />
          </button>
          <div class="-ml-0.5">
            <a href="#" class="-m-1.5 block p-1.5">
              <span class="sr-only">Your Company</span>
              <img class="h-8 w-auto" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600" alt="" />
            </a>
          </div>
        </div>
      </DialogPanel>
    </Dialog>
  </header>

  <div class="mx-auto max-w-6xl pt-8 lg:flex lg:gap-x-16 lg:px-8">
    <main class="px-4 py-16 sm:px-6 lg:flex-auto lg:px-0 lg:py-20">
      <slot />
    </main>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Dialog, DialogPanel } from '@headlessui/vue'
import { Bars3Icon } from '@heroicons/vue/20/solid'
import { XMarkIcon } from '@heroicons/vue/24/outline';
import {useAuthStore} from "@/store/auth";
import router from "@/router";

const mobileMenuOpen = ref(false);

async function logout() {
    try {
      await useAuthStore().logout();
    } catch (e) {
        //
    }

    await router.push('/login');
}
</script>