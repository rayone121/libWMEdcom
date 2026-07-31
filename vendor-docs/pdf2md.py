#!/usr/bin/env python3
"""Convert WinMENTOR documentation PDFs to markdown.

The structure documents are two-column tables — "Parametrul din fisier" on the
left, "Explicatii" on the right — which pdftotext -layout preserves as fixed
columns. The split point is wherever the header word "Explica" starts, so it is
read per document rather than assumed; documents with no such header fall back
to plain prose extraction.
"""
import re, subprocess, sys, unicodedata
from pathlib import Path

FURNITURE = re.compile(
    r'^\s*(?:\d+\s*)?(?:Clasificare:\s*Public|Rev\.\s*\d|winmentor\.ro|'
    r'\.\.\. pentru calculatoare care nu (?:s|ș)tiu contabilitate|'
    r'Creat:\s*\d|Pagina\s*\d|\d+\s*/\s*\d+)\s*.*$', re.I)

def clean(s):
    s = unicodedata.normalize('NFC', s)
    return s.replace('﻿', '').replace('\x00', '').rstrip()

def split_col(lines):
    """Column at which the explanation column starts, or None.

    Derived from the data, not from the "Explicatii" header: the header is
    centred over its column and sits well to the right of where the text
    actually begins. Instead take every line that has content, a run of 3+
    spaces, then more content, and use the most common position at which the
    right-hand run starts.
    """
    from collections import Counter
    starts = Counter()
    for ln in lines:
        for m in re.finditer(r'\S {3,}(?=\S)', ln):
            pos = m.end()
            if 25 < pos < 110:
                starts[pos] += 1
    if not starts:
        return None
    best, n = starts.most_common(1)[0]
    if n < 3:
        return None
    # Snap to the leftmost column within 4 chars that is nearly as common, so a
    # few deeply-indented bullets do not drag the boundary right.
    for p in range(best - 4, best + 1):
        if starts.get(p, 0) >= n * 0.5:
            return p
    return best

def to_md(pdf: Path) -> str:
    raw = subprocess.run(['pdftotext', '-layout', str(pdf), '-'],
                         capture_output=True, text=True, errors='replace').stdout
    lines = [clean(l) for l in raw.split('\n')]
    lines = [l for l in lines if not FURNITURE.match(l)]

    title = next((l.strip() for l in lines if l.strip()), pdf.stem)
    out = [f'# {title}', '',
           f'<sub>Source: `{pdf.name}` — WinMENTOR official documentation, '
           f'converted from PDF.</sub>', '']

    col = split_col(lines)
    if col is None:
        body = '\n'.join(lines)
        body = re.sub(r'\n{3,}', '\n\n', body).strip()
        return '\n'.join(out) + body + '\n'

    rows, cur, orphans, prose = [], None, [], []
    for ln in lines:
        if not ln.strip():
            continue
        # Split at this line's own gap when it has one near the global column;
        # pdftotext pads each row independently, so a fixed cut clips words.
        cut = col
        best = None
        for m in re.finditer(r'\S {3,}(?=\S)', ln):
            if abs(m.end() - col) <= 8 and (best is None or abs(m.end() - col) < abs(best - col)):
                best = m.end()
        if best is not None:
            cut = best
        elif len(ln.rstrip()) > col and ln[col - 1:col + 1].strip():
            # Text runs straight through the column boundary with no gap: this
            # is a full-width prose paragraph, not a table row. Cutting it here
            # would slice a word in half.
            prose.append(ln.strip())
            continue
        left, right = ln[:cut].strip(), ln[cut:].strip()

        if re.match(r'^(Parametrul|Explica)', left):
            continue
        if left:
            if cur:
                rows.append(cur)
            # Explanation blocks are laid out ABOVE the parameter they annotate,
            # so anything collected since the last parameter belongs to this one.
            cur = [left, orphans + ([right] if right else [])]
            orphans = []
        elif right:
            orphans.append(right)
    if cur:
        rows.append(cur)
    if orphans:
        rows.append(['', orphans])

    out += ['| Parametru | Explicație |', '|---|---|']
    for left, expl in rows:
        e = re.sub(r'\s{2,}', ' ', ' '.join(expl).strip()).replace('|', '\\|')
        left = re.sub(r'\s{2,}', ' ', left).replace('|', '\\|')
        out.append(f'| `{left}` | {e} |' if left else f'|  | {e} |')
    if prose:
        out += [''] + [re.sub(r'\s{2,}', ' ', p) for p in prose]
    return '\n'.join(out) + '\n'

if __name__ == '__main__':
    src, dst = Path(sys.argv[1]), Path(sys.argv[2])
    dst.mkdir(parents=True, exist_ok=True)
    n = 0
    for pdf in sorted(src.glob('*.pdf')):
        md = to_md(pdf)
        (dst / (pdf.stem + '.md')).write_text(md, encoding='utf-8')
        n += 1
    print(f'converted {n} PDFs -> {dst}')
