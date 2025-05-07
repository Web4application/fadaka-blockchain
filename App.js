import React, { useState } from "react";
import { IntlProvider, FormattedMessage } from "react-intl";
import Header from "./components/Header";
import "./assets/styles.css";
import logo from "./assets/logo.png";

// Import language files
import en from "./i18n/en.json";
import fr from "./i18n/fr.json";
import es from "./i18n/es.json";

const messages = { en, fr, es };

function App() {
  const [locale, setLocale] = useState("en");

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
        <Header balance={1000} />
      </div>
    </IntlProvider>
  );
}

export default App;
