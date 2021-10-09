#include "vec.h"


data_t dotproduct(vec_ptr u, vec_ptr v) {
   data_t *udata = u->data;
   data_t *vdata = v->data;
   //data_t sum = 0, u_val, v_val;
   data_t sum = 0;
   long n = vec_length(u);
   for (long i = 0; i < n; i++) { // we can assume both vectors are same length
        // get_vec_element(u, i, &u_val);
        //get_vec_element(v, i, &v_val);
        //sum += u_val * v_val;
        sum += udata[i] * vdata[i];
   }   
   return sum;
}
