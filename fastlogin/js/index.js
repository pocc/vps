/* Minimal javascript to switch between tabs */
(() => {
  if (window.location.href.includes('#features')) {
    switchTab('features_tab');
  }
  const tabs = document.getElementsByClassName('nav_tab');
  for (const tab of tabs) {
    tab.addEventListener('click', (e) => {
      const tabLinkId = e.target.id;
      switchTab(tabLinkId);
    });
  }
  function switchTab(tabId) {
    deleteClass('active_tab_link');
    deleteClass('active_tab');
    const tabLink = document.getElementById(tabId);
    tabLink.classList.add('active_tab_link');
    const tabElemId = tabLink.getAttribute('for');
    const tabElem = document.getElementById(tabElemId);
    tabElem.classList.add('active_tab');
  }
  function deleteClass(className) {
    for (const elem of document.getElementsByClassName(className)) {
      elem.classList.remove(className);
    }
  }
})();
/* Minimal javascript to change CSS for slider */
(() => {
  const monthAmt = 1.25;
  const annualAmt = 0.94;
  function getInt(num) {
    return parseInt(num).toString()
  }
  function getMantissa(num) { // 1.23 => "23"
    return (num-parseInt(num)).toFixed(2).substr(2,);
  }
  // Set to annual ammount by default
  document.getElementById('pro_price_int').innerText = getInt(annualAmt);
  document.getElementById('pro_price_decimal').innerText = getMantissa(annualAmt);

  document.getElementById('billing_checkbox').addEventListener('change', (e) => {
    if (e.target.checked) {
      document.getElementById('annual_billing').classList.add('active_billing');
      document.getElementById('monthly_billing').classList.remove('active_billing');
      document.getElementById('pro_price_int').innerText = getInt(annualAmt);
      document.getElementById('pro_price_decimal').innerText = getMantissa(annualAmt);
    } else {
      document.getElementById('monthly_billing').classList.add('active_billing');
      document.getElementById('annual_billing').classList.remove('active_billing');
      document.getElementById('pro_price_int').innerText = getInt(monthAmt);
      document.getElementById('pro_price_decimal').innerText = getMantissa(monthAmt);
    }
  });
})();