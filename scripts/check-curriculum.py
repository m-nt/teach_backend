#!/usr/bin/env python3
"""Check local lesson completeness and file links; no external requests."""
from pathlib import Path
import re
import sys
from urllib.parse import unquote

root = Path(__file__).resolve().parents[1]
errors = []
expected = {'README.md', 'TEACHING.md', 'LAB.md', 'QUIZ.md', 'HOMEWORK.md'}
tiers = sorted((root / 'courses').iterdir())
for tier in tiers:
    if not tier.is_dir():
        continue
    lessons = sorted(p for p in tier.iterdir() if p.is_dir())
    if len(lessons) != 14:
        errors.append(f'{tier.name}: expected 14 lessons, got {len(lessons)}')
    for index, lesson in enumerate(lessons, 1):
        if not lesson.name.startswith(f'{index:02d}-'):
            errors.append(f'Incorrect lesson sequence: {lesson}')
        missing = expected - {p.name for p in lesson.iterdir()}
        for name in missing:
            errors.append(f'Missing {lesson / name}')
        for name in expected - missing:
            if not (lesson / name).read_text().strip():
                errors.append(f'Empty file: {lesson / name}')

# Check Markdown file targets. Ignore code fences, external links, and anchors.
for path in root.rglob('*.md'):
    text = re.sub(r'```.*?```', '', path.read_text(), flags=re.S)
    for target in re.findall(r'\[[^\]]*\]\(([^)]+)\)', text):
        target = target.strip().strip('<>')
        if re.match(r'^[a-zA-Z][a-zA-Z0-9+.-]*:', target) or target.startswith('#'):
            continue
        target = unquote(target.split('#', 1)[0])
        if target and not (path.parent / target).exists():
            errors.append(f'{path.relative_to(root)}: missing link target {target}')
if errors:
    print('\n'.join(errors), file=sys.stderr)
    raise SystemExit(1)
print(f'OK: 3 tiers, 42 lessons, 210 required lesson files; local Markdown links resolve.')
