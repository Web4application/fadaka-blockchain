import React, { useState, useEffect } from "react";
import { IntlProvider, FormattedMessage } from "react-intl";
import Header from "./components/Header";
import Transaction from "./components/Transaction";
import Wallet from "./components/Wallet";
import "./assets/styles.css";
import logo from "./assets/logo.png";

// Import language files
import en from "./i18n/en.json";
import fr from "./i18n/fr.json";
import es from "./i18n/es.json";

const messages = { en, fr, es };

function App() {
  const browserLanguage = navigator.language.slice(0, 2);
  const defaultLanguage = ["en", "fr", "es"].includes(browserLanguage) ? browserLanguage : "en";
  const [locale, setLocale] = useState(defaultLanguage);

  const handleLanguageChange = (e) => {
    setLocale(e.target.value);
  };

  return (
    <IntlProvider locale={locale} messages={messages[locale]}>
      <div className="App">
        <header className="header">
          <img src={logo} className="logo" alt="Fadaka Logo" />
          <h1>
            <FormattedMessage id="app.title" />
          </h1>
          <select value={locale} onChange={handleLanguageChange}>
            <option value="en">English</option>
            <option value="fr">Français</option>
            <option value="es">Español</option>
          </select>
        </header>

        <Wallet address="0x123ABC" balance={5000} />
        <Transaction sender="0x123ABC" receiver="0x456DEF" amount={150} />
      </div>
    </IntlProvider>
  );
}

export default App;
