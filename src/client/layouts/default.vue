<template>
  <div>
    <Navbar
      :company-info="$store.state.company.info"
      :menu-items="$store.state.menu.items"
    />
    <Nuxt />
    <footer class="border-top footer text-muted mt-3 pt-2">
      <div class="container">
        <div class="text-center">
          &copy; Алеф, {{ new Date().getFullYear() }} г. Все права защищены.
        </div>
        <div class="row justify-content-center mt-1">
          <a class="text-muted footer-legal-link" href="/legal/privacy"
            >Политика конфиденциальности</a
          >
          <a class="text-muted footer-legal-link" href="/legal/cookies"
            >Использование cookies</a
          >
        </div>
        <div class="row justify-content-center mt-1">
          <a class="pr-2" href="https://vk.com/alefdoc" target="_blank">
            <ion-icon name="logo-vk" size="small"></ion-icon>
          </a>
          <a
            class="pr-2"
            href="https://www.facebook.com/alefdoc/"
            target="_blank"
          >
            <ion-icon name="logo-facebook" size="small"></ion-icon>
          </a>
          <a
            class="pr-2"
            href="https://www.instagram.com/alefdoc/"
            target="_blank"
          >
            <ion-icon name="logo-instagram" size="small"></ion-icon>
          </a>
          <a
            class="pr-2"
            href="https://zen.yandex.ru/id/5e515cdcfc020165b28c7f75/"
            target="_blank"
          >
            <ion-icon
              src="/images/icon/socials/zen-icon.svg"
              size="small"
            ></ion-icon>
          </a>
        </div>
      </div>
    </footer>
    <script src="https://unpkg.com/ionicons@5.0.0/dist/ionicons.js"></script>
  </div>
</template>

<script>
export default {}
</script>

<style>
html {
  font-family: 'Source Sans Pro', -apple-system, BlinkMacSystemFont, 'Segoe UI',
    Roboto, 'Helvetica Neue', Arial, sans-serif;
  font-size: 16px;
  word-spacing: 1px;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  box-sizing: border-box;
}

*,
*::before,
*::after {
  box-sizing: border-box;
  margin: 0;
}

.button--green {
  display: inline-block;
  border-radius: 4px;
  border: 1px solid #3b8070;
  color: #3b8070;
  text-decoration: none;
  padding: 10px 30px;
}

.button--green:hover {
  color: #fff;
  background-color: #3b8070;
}

.button--grey {
  display: inline-block;
  border-radius: 4px;
  border: 1px solid #35495e;
  color: #35495e;
  text-decoration: none;
  padding: 10px 30px;
  margin-left: 15px;
}

.button--grey:hover {
  color: #fff;
  background-color: #35495e;
}

.footer {
  bottom: 0;
  width: 100%;
  white-space: nowrap;
  padding-top: 10px;
  padding-bottom: 10px;
  font-size: 0.8rem;
}

.footer-legal-link {
  border-right: 1px solid #d6d6d6;
  margin-right: 7px;
  padding-right: 10px;
  display: inline-block;
  white-space: nowrap;
}

.footer-legal-link:last-child {
  border: none;
  margin: 0;
  padding: 0;
}

.footer .row {
  margin-left: 0px;
  margin-right: 0px;
}
</style>
