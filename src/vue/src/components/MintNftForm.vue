<template>
  <div class="sp-type-form sp-box sp-shadow">
    <form class="sp-type-form__main__form">
      <div class="sp-type-form__header sp-box-header">Create an NFT</div>

      <div class="sp-type-form__field sp-form-group">
        <input type="text" class="sp-input" v-model="serial_number_hash" placeholder="Serial Number" required />
      </div>

      <div class="sp-type-form__field sp-form-group">
        <input type="text" class="sp-input" v-model="manufacturer" placeholder="Manufacturer" required />
      </div>

      <div class="sp-type-form__field sp-form-group">
        <input type="text" class="sp-input" v-model="part_number" placeholder="Part Number" required />
      </div>

      <div class="sp-type-form__field sp-form-group">
        <input type="text" class="sp-input" v-model="uri" placeholder="Metadata URI" required />
      </div>

      <div class="sp-type-form__field sp-form-group">
        <input type="text" class="sp-input" placeholder="Metadata URI Hash" v-model="uri_hash" required />
      </div>

      <div class="sp-type-form__field sp-form-group">
        <input type="text" class="sp-input" v-model="trust_anchor_token" placeholder="Physical Asset Owner ID" />
      </div>
      <SpButton type="secondary" v-on:click="getToken">Get Token</SpButton>

      <div class="sp-type-form__field sp-form-group">
        <input type="text" class="sp-input" v-model="trust_anchor" placeholder="Trust Anchor (Registered Agent) ID" disabled />
      </div>

      <div class="sp-type-form__btns">
        <SpButton type="primary" v-on:click="submit">Mint OBT</SpButton>
      </div>
    </form>
  </div>
</template>

<style>
.sp-type-form__btns {
  margin-top: 2rem;
}
</style>

<script>
import axios from 'axios'

export default {
  name: "MintNftForm",
  data() {
    return {
      // Until we not define trust-anchor spec with Denis and Joe we going to use sandbox service
      taAuthToken: "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZWU0MjRmN2QtODVjNi00MTM2LWJmMjQtYTgxYjA3ZGI2OTY5In0.3sl5tmp_BOmHSRpQbWMSK5CpDmHQLNOHpy1WzNgrkvDbEi4P3DdpiD-1iH9QsQajx_JZhkRq3S4-Ti1Cs390Bw",
      trustAnchors: [{
        id: "demoTrustAnchor",
        name: "Demo Trust Anchor"
      }],
      serial_number_hash: "",
      manufacturer: "",
      part_number: "",
      trust_anchor_token: "",
      uri: "",
      uri_hash: "",
      trust_anchor: "demoTrustAnchor",
    };
  },
  computed: {

    currentAccount() {
      if (this._depsLoaded) {
        if (this.loggedIn) {
          return this.$store.getters['common/wallet/address']
        } else {
          return null
        }
      } else {
        return null
      }
    },
    loggedIn() {
      if (this._depsLoaded) {
        return this.$store.getters['common/wallet/loggedIn']
      } else {
        return false
      }
    }
  },
  methods: {
    getToken() {
      axios.post("http://demo.ta.alpha.obada.io/api/v1/issue-token", {}, {
        headers: {
          'Authorization': 'Bearer ' + this.taAuthToken,
          'Content-Type': 'application/json'
        },
        withCredentials: false
      })
          .then(response => {
            this.trust_anchor_token = response.data.token
          })
          .catch(error => {
            this.errorMessage = error.message;
            console.error("There was an error!", error);
          });

      return ""
    },
    async submit() {
      const value = {
        creator: this.currentAccount,
        serialNumberHash: this.serial_number_hash,
        manufacturer: this.manufacturer,
        partNumber: this.part_number,
        trustAnchorToken: this.trust_anchor_token,
        uri: this.uri,
        uriHash: this.uri_hash,
      };

      const resp = await this.$store.dispatch("obadafoundation.fullcore.obit/sendMsgMintObit", {
        value,
        fee: [],
      });

      console.log(resp);
    },
  },
};
</script>
