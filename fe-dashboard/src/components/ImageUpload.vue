<template>
        <div class="upload">
            
            <div class="image-preview" v-if="imageData.length > 0">
                <div v-for="img, index in images" :key="index">
                    <img class="preview" :src="img">
                </div>
            </div>
            <input type="file" ref="file" @change="previewImage" accept="image/*">
            <button class="btn btn-light" @click="$refs.file.click()">+</button>
        </div>
</template>
<script>

export default {
    name: "ImageUpload",
    data() {
        return{
            imageData: "",
            file:'',
            images: [],
        }
        
    },
    methods: {
        previewImage(event) {
            this.file = this.$refs.file.files[0];
            console.log(this.file)
            this.$emit("setFile", this.file);
            
            let input = event.target;
            if (input.files && input.files[0]) {
                let reader = new FileReader();
                reader.onload = (e) => {
                    this.imageData = e.target.result;
                    this.images.push(this.imageData);
                }
                reader.readAsDataURL(input.files[0]);
            }
        },
    }
}
</script>

<style scoped>
.upload {
    display: flex;
    flex-direction: row;
}

.image-preview {
    display: flex;
}

img.preview {
    width: 100px;
    height: 100px;
    background-color: white;
    border: 1px solid #DDD;
    padding: 2px;
    margin-right: 5px;
}

input {
    display: none;
}

button {
    height: 100px;
    width: 100px;
    background-color: #fff;
    border: 1px solid #7f8c8d;
    border-style: dashed;
    cursor: pointer;
    margin-bottom: 30px;
}

</style>