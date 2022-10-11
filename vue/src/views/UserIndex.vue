<template>
    <CardContainer>
        User
    </CardContainer>
</template>

<script>
import CardContainer from '@/components/CardContainer'
import $ from 'jquery'

export default {
    components: {
        CardContainer,
    },
    setup() {
        $.ajax({
            url: "http://127.0.0.1:3000/api/",
            type: "get",
            success(resp) {
                console.log(resp);
            },
            error(resp) {
                console.log(resp);
            }
        });
    }
}
</script>

<style scoped></style>