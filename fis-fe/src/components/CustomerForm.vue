<template>
<div class="modal fade" id="signin-modal" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered" role="document">
            <div class="modal-content">
                <div class="modal-body">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                        <span aria-hidden="true"><i class="icon-close"></i></span>
                    </button>

                    <div class="form-box">
                        <div class="form-tab">
                            <ul class="nav nav-pills nav-fill" role="tablist">
                                <li class="nav-item">
                                    <a class="nav-link active" id="signin-tab" data-toggle="tab" href="#signin" role="tab" aria-controls="signin" aria-selected="true">Đăng ký tư vấn</a>
                                </li>
                            </ul>
                            <div class="tab-content" id="tab-content-5">
                                <div class="tab-pane fade show active" id="signin" role="tabpanel" aria-labelledby="signin-tab">
                                    <form @submit.prevent="createCustomer">
                                        <div class="form-group">
                                            <label>Họ và tên *</label>
                                            <input type="text" class="form-control" name="full_name" v-model="name" required>
                                        </div><!-- End .col-sm-6 -->
                                        <div class="form-group">
                                            <label >Số điện thoại *</label>
                                            <input type="tel" class="form-control" name="phone_number" v-model="phone_number" required>
                                        </div><!-- End .col-sm-6 -->
                                        <div class="form-group">
                                            <label>Email *</label>
                                            <input type="text" class="form-control"  name="email" v-model="email" required>
                                        </div><!-- End .form-group -->
                                        <div class="form-group">
                                            <label>Sản phẩm *</label>
                                            <input type="text" class="form-control" name="product" v-model="product_info" required>
                                        </div><!-- End .form-group -->
                                        <div class="form-group">
                                            <label for="cmessage">Nội dung chi tiết</label>
                                            <textarea class="form-control" cols="30" rows="4" name="message" v-model="message"></textarea>
                                        </div>
                                        <div class="form-footer">
                                            <button type="submit" class="btn btn-outline-primary-2">
                                                <span>Đăng ký tư vấn</span>
                                                <i class="icon-long-arrow-right"></i>
                                            </button>
                                        </div><!-- End .form-footer -->
                                    </form>
                                </div><!-- .End .tab-pane -->
                            </div><!-- End .tab-content -->
                        </div><!-- End .form-tab -->
                    </div><!-- End .form-box -->
                </div><!-- End .modal-body -->
            </div><!-- End .modal-content -->
        </div><!-- End .modal-dialog -->
    </div><!-- End .modal -->
</template>
<script>
export default {
    name : 'CustomerForm',
    // props: ["name"],
    data() {
        return {
            name : "",
            email: "",
            phone_number: "",
            product_info: "",
            subject: "",
            message: "",
        };
    },
    methods: {
        async createCustomer() {
            const customer = {
                name : this.name,
                email : this.email,
                phone_number : this.phone_number,
                product_info: this.product_info,
                subject : this.subject,
                message : this.message,
            }
            await this.$store.dispatch("customers/createCustomer", customer).then(
                alert('Đăng ký thành công')
            )
            window.location.reload();
        },
    }
}
</script>
