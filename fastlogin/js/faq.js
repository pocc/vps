/* Easy way to load html */
const faqText = `<h1 id="login-faster-faq" class="title_size">FAQ</h1>` + 
`<h2 id="what-does-this-extension-do">What does this extension do?</h2>
<ul>
<li>Decrease your Meraki login time by a factor of 10</li>
<li>Quickly switch between accounts</li>
<li>Any other minor UI tweaks that users want.</li>
</ul>
<h2 id="why-are-accounts-highlighted-different-colors-and-what-does-the-key-mean">Why are accounts highlighted different colors? And what does the key mean?</h2>
<ul>
<li><strong>key</strong>: Account is currently logged in</li>
<li><strong>black</strong>: This account is valid and there has been a successful login attempt</li>
<li><strong>blue</strong>: Account is logged in and you can click on the link to open the last visited page.</li>
<li><strong>pink</strong>: Account has bad user/pass. You will need to reenter credentials.</li>
</ul>
<h2 id="whats-the-difference-between-clicking-on-an-account-link-and-the-login-button">What's the difference between clicking on an account link and the login button?</h2>
<p>If you are already logged in, a link will be added to the account to the last url that that account visited. If you click on the login button, you will login to that account. If you are already logged in, you will be logged out and login again.</p>
<h2 id="how-does-relogin-collect-meraki-credentials">How does Relogin collect Meraki credentials?</h2>
<ul>
<li>You enter them manually in the popup</li>
<li>You enter them on Meraki's login page</li>
</ul>
<h2 id="i-found-a-bug-what-do-i-do">I found a bug what do I do?</h2>
<p>Report it with this <a href="https://docs.google.com/forms/d/1CPFoAolM3TssRo3rebbmYauoEDuuFe_BQ64_7ucYTYg">google form</a>. Note that you might be able to fix it by refreshing the page.</p>
<h2 id="how-does-relogin-store-meraki-credentials">How does Relogin store Meraki credentials?</h2>
<p>Firefox and Chrome will both offer to save passwords when they see that you have entered one.
Login Faster stores Meraki credentials on Google's/Firefox's servers in a similar manner.</p>
<h2 id="what-data-collected-about-me">What data collected about me?</h2>
<p>Google Tag Manager collects anonymized usage statistics.</p>`
document.getElementById('faq_tab').innerHTML = faqText;