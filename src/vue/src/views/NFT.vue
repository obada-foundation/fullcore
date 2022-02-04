<template>
  <div>
    <div class="container">
      <h3>OBT NFT</h3>
      <MintNftForm />
    </div>
  </div>
</template>

<script>
import MintNftForm from "../components/MintNftForm";
  
export default {
  name: 'NFT',
  components: { MintNftForm },
}
</script>
