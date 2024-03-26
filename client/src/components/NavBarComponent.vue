<template>
  <b-navbar type="dark" variant="info" class="fixed-top">
    <b-navbar-nav>
      <b-nav-item href="/">Home</b-nav-item>
      <b-nav-item href="/about">About</b-nav-item>
      <b-nav-item href="/results">Results</b-nav-item>
    </b-navbar-nav>
    <b-navbar-nav class="ml-auto">
      <b-nav-item-dropdown text="EN" right>
        <b-dropdown-item href="#">ES</b-dropdown-item>
        <b-dropdown-item href="#">FR</b-dropdown-item>
      </b-nav-item-dropdown>
      <b-nav-item-dropdown right>
        <template #button-content>User</template>
        <b-dropdown-item href="#">Profile</b-dropdown-item>
        <b-dropdown-item href="#">Sign Out</b-dropdown-item>
      </b-nav-item-dropdown>
    </b-navbar-nav>
  </b-navbar>
</template>

<style>
.navbar-dark .navbar-nav .nav-link {
  color: white !important;
}
</style>
