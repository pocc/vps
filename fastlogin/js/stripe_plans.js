/* From https://dashboard.stripe.com/subscriptions/products/ for monthly recurring plan */
(() => {
  const stripe = Stripe('pk_live_dCgtCLRtnCed6YB98Lu0Y8mh001hytA5CX');
  
  const checkoutButton = document.getElementById('checkout-button');
  checkoutButton.addEventListener('click', function() {
    // When the customer clicks on the button, redirect
    // them to Checkout.
    const successUrl = 'https://fastlog.in/#success';
    const cancelUrl = 'https://fastlog.in';

    const annualPlan = 'plan_GirSsO0M7UgjKj'; 
    const monthlyPlan = 'plan_GirXrr9RnpO82e';
    const plan = document.getElementById('billing_checkbox').checked ? annualPlan : monthlyPlan;
    const data = {
      items: [{plan: plan, quantity: 1}],

      // Do not rely on the redirect to the successUrl for fulfilling
      // purchases, customers may not always reach the success_url after
      // a successful payment.
      // Instead use one of the strategies described in
      // https://stripe.com/docs/payments/checkout/fulfillment
      successUrl: successUrl,
      cancelUrl: cancelUrl,
    };
    // If an email has been passed in via the url, pass it along to stripe
    if (window.location.hash.includes('?')) {
      const path = '?' + window.location.hash.split('?')[1]; 
      const emailCandidate = new URLSearchParams(path).get('email')
      if (emailCandidate !== null && emailCandidate.length > 0) {
        data.customerEmail = emailCandidate; 
      }
    }
    stripe.redirectToCheckout(data)
    .then(function (result) {
      if (result.error) {
        // If `redirectToCheckout` fails due to a browser or network
        // error, display the localized error message to your customer.
        var displayError = document.getElementById('error-message');
        displayError.textContent = result.error.message;
      }
    });
  });
})();
