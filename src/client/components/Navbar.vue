<template>
  <b-navbar toggleable="lg" type="light" variant="faded">
    <div class="container">
      <router-link class="navbar-brand" to="/">
        <img
          class="mx-auto d-block"
          height="32"
          src="/logo_with_text.svg"
          alt="alefdoc"
        />
      </router-link>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav class="mx-auto">
          <b-nav-item
            v-for="(item, i) in menuItems"
            :key="i"
            :to="item.url"
            class="navbar-light font-weight-bold"
          >
            {{ item.title }}
          </b-nav-item>
        </b-navbar-nav>

        <b-navbar-nav class="ml-auto">
          <span class="navbar-text pt-0 pb-0 text-center">
            <span class="navbar-text p-0 mr-3 text-center">
              <small>Контактный телефон</small>
              <a
                class="nav-link text-dark font-weight-bold p-0"
                :href="'tel:' + companyInfo.phone"
                >{{ companyInfo.phone }}
              </a>
            </span>
            <span class="navbar-text p-0 mr-0 text-center">
              <small>Электронная почта</small>
              <a
                class="nav-link text-dark font-weight-bold p-0"
                :href="'mailto:' + companyInfo.email"
                >{{ companyInfo.email }}
              </a>
            </span>
          </span>
        </b-navbar-nav>
      </b-collapse>
    </div>
  </b-navbar>
</template>

<script>
export default {
  props: {
    companyInfo: {
      type: Object,
      default() {
        return null
      },
    },
    menuItems: {
      type: Array,
      default() {
        return []
      },
    },
  },
}
</script>
