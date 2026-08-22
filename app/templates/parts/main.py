import os
import re

__call__ = lambda fn: print(fn())


def extract_css_classes_from_file(filepath):
    classes: set = set()

    with open(filepath, 'r', encoding='utf-8') as f:
        content = f.read()

    # Busca todos los atributos class="..."
    class_pattern = r'class=["\']([^"\']*)["\']'
    matches: list = re.findall(class_pattern, content)

    for match in matches:
        # Divide por espacios y filtra clases vacías
        for cls in match.split():
            if cls.strip():
                classes.add(cls.strip())

    return classes


def extract_css_names_from_selectors(filepath):
    selectors = []

    with open(filepath, 'r', encoding='utf-8') as f:
        for line in f:
            selector = line.strip()
            if not selector:
                continue

            # Limpia el selector
            # Elimina . al inicio y :hover, etc.
            cleaned = selector.replace('.', '').split(':')[0].split()[0]
            if cleaned:
                selectors.append(cleaned)

    return list(set(selectors))

def main():
    # Extrae clases CSS de los selectores
    css_selectors = extract_css_names_from_selectors('./classess')
    print(f"Total de selectores CSS: {len(css_selectors)}")

    # Verifica cada archivo HTML
    html_files = [f for f in os.listdir() if f.endswith('.html')]
    all_found = set()

    for html_file in html_files:
        html_classes = extract_css_classes_from_file(html_file)
        print(f"\n{html_file}:")
        print(f"  Clases en HTML: {len(html_classes)}")

        # Encuentra qué selectores CSS están presentes
        found = set()
        missing = []

        for selector in css_selectors:
            if selector in html_classes:
                found.add(selector)
            else:
                missing.append(selector)

        print(f"  Selectores encontrados: {len(found)}")
        all_found.update(found)

        if missing:
            print(f"  Selectores faltantes ({len(missing)}):")
            for m in missing[:10]:  # Muestra primeros 10
                print(f"    - {m}")
            if len(missing) > 10:
                print(f"    ... y {len(missing) - 10} más")

    print(f"\nRESUMEN FINAL:")
    print(f"Total selectores CSS: {len(css_selectors)}")
    print(f"Total selectores encontrados en todos los HTMLs: {len(all_found)}")
    print(f"Total selectores NO encontrados: {len(css_selectors) - len(all_found)}")

    # Muestra los no encontrados
    missing_all = set(css_selectors) - all_found
    if missing_all:
        print(f"\nSelectores NO encontrados en ningún HTML:")
        for m in sorted(missing_all):
            print(f"  - {m}")

if __name__ == "__main__":
    main()
