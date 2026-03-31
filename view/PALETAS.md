# 🎨 Paletas de Cores - Clown CRM

## Temas Disponíveis

### 🍷 Burgundy (Padrão)
```
Primária:    #8B0E0E ████████
Dark:        #5C0A0A ████████
Light:       #C91A1A ████████
Background:  #0A0A0A ████████
Surface:     #1A1A1A ████████
Text:        #FFFFFF ████████
```
**Uso**: Tema elegante e profissional com tons de vinho

---

### 💜 Dark Purple
```
Primária:    #7C3AED ████████
Dark:        #5B21B6 ████████
Light:       #A78BFA ████████
Background:  #0F0A1E ████████
Surface:     #1A1333 ████████
Text:        #C4B5FD ████████
```
**Uso**: Tema místico e criativo, perfeito para aplicações de design

---

### 🌊 Ocean Blue
```
Primária:    #0EA5E9 ████████
Dark:        #0369A1 ████████
Light:       #38BDF8 ████████
Background:  #020617 ████████
Surface:     #0F172A ████████
Text:        #94A3B8 ████████
```
**Uso**: Tema calmo e profissional, ideal para produtividade

---

### 🌲 Forest Green
```
Primária:    #10B981 ████████
Dark:        #059669 ████████
Light:       #34D399 ████████
Background:  #0A1810 ████████
Surface:     #132B23 ████████
Text:        #A7F3D0 ████████
```
**Uso**: Tema natural e refrescante, ótimo para sustentabilidade

---

### 🌅 Sunset Orange
```
Primária:    #F97316 ████████
Dark:        #C2410C ████████
Light:       #FB923C ████████
Background:  #1C0E08 ████████
Surface:     #2A1710 ████████
Text:        #FDBA74 ████████
```
**Uso**: Tema quente e energético, perfeito para call-to-actions

---

### 🌙 Midnight Blue
```
Primária:    #3B82F6 ████████
Dark:        #1E40AF ████████
Light:       #60A5FA ████████
Background:  #030712 ████████
Surface:     #111827 ████████
Text:        #9CA3AF ████████
```
**Uso**: Tema clássico e confiável, universal para qualquer aplicação

---

## Cores Semânticas (Todas as Paletas)

### Success (Sucesso)
```
#10B981 ████████ - Verde
```
**Uso**: Confirmações, sucesso em operações, estados positivos

### Error (Erro)
```
#EF4444 ████████ - Vermelho
```
**Uso**: Erros, alertas críticos, ações destrutivas

### Warning (Aviso)
```
#F59E0B ████████ - Laranja/Amarelo
```
**Uso**: Avisos, ações que requerem atenção

### Info (Informação)
```
#3B82F6 ████████ - Azul
```
**Uso**: Informações neutras, tooltips, ajuda

---

## Exemplo de Uso por Contexto

### Dashboard Corporativo
**Recomendado**: 🌙 Midnight Blue ou 🌊 Ocean Blue
- Profissional e confiável
- Boa legibilidade
- Calmo para uso prolongado

### Aplicação Criativa
**Recomendado**: 💜 Dark Purple ou 🌅 Sunset Orange
- Energético e inspirador
- Destaque visual
- Expressão de personalidade

### Plataforma de Produtividade
**Recomendado**: 🌲 Forest Green ou 🌊 Ocean Blue
- Reduz fadiga visual
- Promove foco
- Ambiente calmo

### Sistema Administrativo
**Recomendado**: 🍷 Burgundy ou 🌙 Midnight Blue
- Profissional e elegante
- Hierarquia visual clara
- Tradicional e confiável

---

## Acessibilidade

Todas as paletas foram projetadas com:
- ✅ Contraste adequado (WCAG AA)
- ✅ Texto legível em todos os fundos
- ✅ Estados de foco visíveis
- ✅ Feedback visual claro

## Personalização

Para criar uma paleta customizada, edite `src/styles/theme.rs`:

```rust
pub fn my_theme() -> Self {
    Theme {
        primary: "#YOUR_PRIMARY",
        primary_dark: "#YOUR_DARK",
        primary_light: "#YOUR_LIGHT",
        background: "#YOUR_BG",
        surface: "#YOUR_SURFACE",
        // ... outras cores
    }
}
```

Adicione ao enum `ThemeVariant`:
```rust
pub enum ThemeVariant {
    // ... existing themes
    MyCustomTheme,
}
```

E ao match em `from_variant`:
```rust
ThemeVariant::MyCustomTheme => Self::my_theme(),
```
