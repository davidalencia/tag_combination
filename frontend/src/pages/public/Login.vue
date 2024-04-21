<script setup>
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';
import InputText from 'primevue/inputtext';
import Password from 'primevue/password';
import Button from 'primevue/button';

</script>

<template>

  <div>
    <InputGroup>
      <InputGroupAddon>
        <i class="pi pi-user"></i>
      </InputGroupAddon>
      <InputText placeholder="Username" />
    </InputGroup>
    <InputGroup>
      <InputGroupAddon>
        <i class="pi pi-key"></i>
      </InputGroupAddon>
      <Password v-model="value" :feedback="false" placeholder="Password" />
    </InputGroup>
    <Button>Login</Button>
  </div>
  <div>
    <InputGroup>
      <InputGroupAddon>
        <i class="pi pi-user"></i>
      </InputGroupAddon>
      <InputText placeholder="Username" />
    </InputGroup>
    <InputGroup>
      <InputGroupAddon>
        <i class="pi pi-key"></i>
      </InputGroupAddon>
      <Password v-model="value" :feedback="false" placeholder="Password" />
    </InputGroup>
    <Button>Register</Button>
  </div>
</template>