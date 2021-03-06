<template>
  <b-navbar toggleable="lg" type="light" variant="faded">
    <div class="container">
      <b-navbar-brand href="#">
        <img
          src="/logo_with_text.svg"
          class="d-inline-block align-top"
          alt="alefdoc"
          height="32"
        />
      </b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav class="mx-auto">
          <b-nav-item
            v-for="(link, i) in links"
            :key="i"
            :to="link.to"
            class="navbar-light font-weight-bold"
          >
            {{ link.title }}
          </b-nav-item>
        </b-navbar-nav>

        <b-navbar-nav class="ml-auto">
          <span class="navbar-text pt-0 pb-0 text-center">
            <span class="navbar-text p-0 mr-3 text-center">
              <small>Контактный телефон</small>
              <a
                class="nav-link text-dark font-weight-bold p-0"
                :href="'tel:' + phone"
                >{{ phone }}</a
              >
            </span>
            <span class="navbar-text p-0 mr-0 text-center">
              <small>Электронная почта</small>
              <a
                class="nav-link text-dark font-weight-bold p-0"
                :href="'mailto:' + email"
                >{{ email }}</a
              >
            </span>
          </span>
        </b-navbar-nav>
      </b-collapse>
    </div>
  </b-navbar>
</template>

<script>
export default {
  async asyncData({ $axios }) {
    const ip = await $axios.$get('http://icanhazip.com')
    return { ip }
  },
  data() {
    return {
      phone: '+7(789)123-45-67',
      email: 'info@example.ru',
      links: [
        { to: '#', title: 'Ссылка 1' },
        { to: '#', title: 'Ссылка 2' },
        { to: '#', title: 'Ссылка 3' },
        { to: '#', title: 'Ссылка 4' },
        { to: '#', title: 'Ссылка 5' },
        { to: '#', title: 'Ссылка 6' },
      ],
    }
  },
}
</script>
