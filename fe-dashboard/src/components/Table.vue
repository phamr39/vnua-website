<template>
    <table class="table table-bordered">
        <thead>
            <tr>
                <th scope="col">Họ tên</th>
                <th scope="col">Email</th>
                <th scope="col">Điện thoại</th>
                <th scope="col">Tin nhắn</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="post in posts" :key="post.id">
                <td >{{post.name}}</td>
                <td >{{post.email}}</td>
                <td >{{post.phone_number}}</td>
                <td >{{post.message}}</td>
            </tr>
        </tbody>
    </table>
</template>
<script>
export default {
    name : 'Table',
    props : ["posts"],
}
</script>
